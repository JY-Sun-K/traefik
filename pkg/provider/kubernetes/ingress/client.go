package ingress

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/rs/zerolog/log"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	kerror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kinformers "k8s.io/client-go/informers"
	kclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"traefik/v3/pkg/provider/kubernetes/k8s"
	traefikversion "traefik/v3/pkg/version"
)

const (
	resyncPeriod   = 10 * time.Minute
	defaultTimeout = 5 * time.Second
)

// Client is a client for the Provider master.
// WatchAll starts the watch of the Provider resources and updates the stores.
// The stores can then be accessed via the Get* functions.
type Client interface {
	WatchAll(namespaces []string, stopCh <-chan struct{}) (<-chan interface{}, error)
	GetIngresses() []*netv1.Ingress
	GetIngressClasses() ([]*netv1.IngressClass, error)
	GetService(namespace, name string) (*corev1.Service, bool, error)
	GetSecret(namespace, name string) (*corev1.Secret, bool, error)
	GetEndpoints(namespace, name string) (*corev1.Endpoints, bool, error)
	UpdateIngressStatus(ing *netv1.Ingress, ingStatus []netv1.IngressLoadBalancerIngress) error
}

type clientWrapper struct {
	clientset                   kclientset.Interface
	factoriesKube               map[string]kinformers.SharedInformerFactory
	factoriesSecret             map[string]kinformers.SharedInformerFactory
	factoriesIngress            map[string]kinformers.SharedInformerFactory
	clusterFactory              kinformers.SharedInformerFactory
	ingressLabelSelector        string
	isNamespaceAll              bool
	disableIngressClassInformer bool
	watchedNamespaces           []string
	serverVersion               *version.Version
}

// newInClusterClient returns a new Provider client that is expected to run
// inside the cluster.
func newInClusterClient(endpoint string) (*clientWrapper, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create in-cluster configuration: %w", err)
	}

	if endpoint != "" {
		config.Host = endpoint
	}

	return createClientFromConfig(config)
}

func newExternalClusterClientFromFile(file string) (*clientWrapper, error) {
	configFromFlags, err := clientcmd.BuildConfigFromFlags("", file)
	if err != nil {
		return nil, err
	}
	return createClientFromConfig(configFromFlags)
}

// newExternalClusterClient returns a new Provider client that may run outside
// of the cluster.
// The endpoint parameter must not be empty.
func newExternalClusterClient(endpoint, token, caFilePath string) (*clientWrapper, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint missing for external cluster client")
	}

	config := &rest.Config{
		Host:        endpoint,
		BearerToken: token,
	}

	if caFilePath != "" {
		caData, err := os.ReadFile(caFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA file %s: %w", caFilePath, err)
		}

		config.TLSClientConfig = rest.TLSClientConfig{CAData: caData}
	}
	return createClientFromConfig(config)
}

func createClientFromConfig(c *rest.Config) (*clientWrapper, error) {
	c.UserAgent = fmt.Sprintf(
		"%s/%s (%s/%s) kubernetes/ingress",
		filepath.Base(os.Args[0]),
		traefikversion.Version,
		runtime.GOOS,
		runtime.GOARCH,
	)

	clientset, err := kclientset.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	return newClientImpl(clientset), nil
}

func newClientImpl(clientset kclientset.Interface) *clientWrapper {
	return &clientWrapper{
		clientset:        clientset,
		factoriesSecret:  make(map[string]kinformers.SharedInformerFactory),
		factoriesIngress: make(map[string]kinformers.SharedInformerFactory),
		factoriesKube:    make(map[string]kinformers.SharedInformerFactory),
	}
}

// WatchAll starts namespace-specific controllers for all relevant kinds.
func (c *clientWrapper) WatchAll(namespaces []string, stopCh <-chan struct{}) (<-chan interface{}, error) {
	// Get and store the serverVersion for future use.
	serverVersionInfo, err := c.clientset.Discovery().ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve server version: %w", err)
	}

	serverVersion, err := version.NewVersion(serverVersionInfo.GitVersion)
	if err != nil {
		return nil, fmt.Errorf("could not parse server version: %w", err)
	}

	c.serverVersion = serverVersion

	eventCh := make(chan interface{}, 1)
	eventHandler := &k8s.ResourceEventHandler{Ev: eventCh}

	if len(namespaces) == 0 {
		namespaces = []string{metav1.NamespaceAll}
		c.isNamespaceAll = true
	}

	c.watchedNamespaces = namespaces

	notOwnedByHelm := func(opts *metav1.ListOptions) {
		opts.LabelSelector = "owner!=helm"
	}

	matchesLabelSelector := func(opts *metav1.ListOptions) {
		opts.LabelSelector = c.ingressLabelSelector
	}

	for _, ns := range namespaces {
		factoryIngress := kinformers.NewSharedInformerFactoryWithOptions(c.clientset, resyncPeriod, kinformers.WithNamespace(ns), kinformers.WithTweakListOptions(matchesLabelSelector))

		_, err = factoryIngress.Networking().V1().Ingresses().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}

		c.factoriesIngress[ns] = factoryIngress

		factoryKube := kinformers.NewSharedInformerFactoryWithOptions(c.clientset, resyncPeriod, kinformers.WithNamespace(ns))
		_, err = factoryKube.Core().V1().Services().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}
		_, err = factoryKube.Core().V1().Endpoints().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}
		c.factoriesKube[ns] = factoryKube

		factorySecret := kinformers.NewSharedInformerFactoryWithOptions(c.clientset, resyncPeriod, kinformers.WithNamespace(ns), kinformers.WithTweakListOptions(notOwnedByHelm))
		_, err = factorySecret.Core().V1().Secrets().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}
		c.factoriesSecret[ns] = factorySecret
	}

	for _, ns := range namespaces {
		c.factoriesIngress[ns].Start(stopCh)
		c.factoriesKube[ns].Start(stopCh)
		c.factoriesSecret[ns].Start(stopCh)
	}

	for _, ns := range namespaces {
		for typ, ok := range c.factoriesIngress[ns].WaitForCacheSync(stopCh) {
			if !ok {
				return nil, fmt.Errorf("timed out waiting for controller caches to sync %s in namespace %q", typ, ns)
			}
		}

		for typ, ok := range c.factoriesKube[ns].WaitForCacheSync(stopCh) {
			if !ok {
				return nil, fmt.Errorf("timed out waiting for controller caches to sync %s in namespace %q", typ, ns)
			}
		}

		for typ, ok := range c.factoriesSecret[ns].WaitForCacheSync(stopCh) {
			if !ok {
				return nil, fmt.Errorf("timed out waiting for controller caches to sync %s in namespace %q", typ, ns)
			}
		}
	}

	if !c.disableIngressClassInformer {
		c.clusterFactory = kinformers.NewSharedInformerFactoryWithOptions(c.clientset, resyncPeriod)

		_, err = c.clusterFactory.Networking().V1().IngressClasses().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}

		c.clusterFactory.Start(stopCh)

		for typ, ok := range c.clusterFactory.WaitForCacheSync(stopCh) {
			if !ok {
				return nil, fmt.Errorf("timed out waiting for controller caches to sync %s", typ)
			}
		}
	}

	return eventCh, nil
}

// GetIngresses returns all Ingresses for observed namespaces in the cluster.
func (c *clientWrapper) GetIngresses() []*netv1.Ingress {
	var results []*netv1.Ingress

	for ns, factory := range c.factoriesIngress {
		// networking
		listNew, err := factory.Networking().V1().Ingresses().Lister().List(labels.Everything())
		if err != nil {
			log.Error().Err(err).Msgf("Failed to list ingresses in namespace %s", ns)
			continue
		}

		results = append(results, listNew...)
	}

	return results
}

// UpdateIngressStatus updates an Ingress with a provided status.
func (c *clientWrapper) UpdateIngressStatus(src *netv1.Ingress, ingStatus []netv1.IngressLoadBalancerIngress) error {
	if !c.isWatchedNamespace(src.Namespace) {
		return fmt.Errorf("failed to get ingress %s/%s: namespace is not within watched namespaces", src.Namespace, src.Name)
	}

	ing, err := c.factoriesIngress[c.lookupNamespace(src.Namespace)].Networking().V1().Ingresses().Lister().Ingresses(src.Namespace).Get(src.Name)
	if err != nil {
		return fmt.Errorf("failed to get ingress %s/%s: %w", src.Namespace, src.Name, err)
	}

	logger := log.With().Str("namespace", ing.Namespace).Str("ingress", ing.Name).Logger()

	if isLoadBalancerIngressEquals(ing.Status.LoadBalancer.Ingress, ingStatus) {
		logger.Debug().Msg("Skipping ingress status update")
		return nil
	}

	ingCopy := ing.DeepCopy()
	ingCopy.Status = netv1.IngressStatus{LoadBalancer: netv1.IngressLoadBalancerStatus{Ingress: ingStatus}}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	_, err = c.clientset.NetworkingV1().Ingresses(ingCopy.Namespace).UpdateStatus(ctx, ingCopy, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update ingress status %s/%s: %w", src.Namespace, src.Name, err)
	}

	logger.Info().Msg("Updated ingress status")
	return nil
}

// isLoadBalancerIngressEquals returns true if the given slices are equal, false otherwise.
func isLoadBalancerIngressEquals(aSlice, bSlice []netv1.IngressLoadBalancerIngress) bool {
	if len(aSlice) != len(bSlice) {
		return false
	}

	aMap := make(map[string]struct{})
	for _, aIngress := range aSlice {
		aMap[aIngress.Hostname+aIngress.IP] = struct{}{}
	}

	for _, bIngress := range bSlice {
		if _, exists := aMap[bIngress.Hostname+bIngress.IP]; !exists {
			return false
		}
	}

	return true
}

// GetService returns the named service from the given namespace.
func (c *clientWrapper) GetService(namespace, name string) (*corev1.Service, bool, error) {
	if !c.isWatchedNamespace(namespace) {
		return nil, false, fmt.Errorf("failed to get service %s/%s: namespace is not within watched namespaces", namespace, name)
	}

	service, err := c.factoriesKube[c.lookupNamespace(namespace)].Core().V1().Services().Lister().Services(namespace).Get(name)
	exist, err := translateNotFoundError(err)
	return service, exist, err
}

// GetEndpoints returns the named endpoints from the given namespace.
func (c *clientWrapper) GetEndpoints(namespace, name string) (*corev1.Endpoints, bool, error) {
	if !c.isWatchedNamespace(namespace) {
		return nil, false, fmt.Errorf("failed to get endpoints %s/%s: namespace is not within watched namespaces", namespace, name)
	}

	endpoint, err := c.factoriesKube[c.lookupNamespace(namespace)].Core().V1().Endpoints().Lister().Endpoints(namespace).Get(name)
	exist, err := translateNotFoundError(err)
	return endpoint, exist, err
}

// GetSecret returns the named secret from the given namespace.
func (c *clientWrapper) GetSecret(namespace, name string) (*corev1.Secret, bool, error) {
	if !c.isWatchedNamespace(namespace) {
		return nil, false, fmt.Errorf("failed to get secret %s/%s: namespace is not within watched namespaces", namespace, name)
	}

	secret, err := c.factoriesSecret[c.lookupNamespace(namespace)].Core().V1().Secrets().Lister().Secrets(namespace).Get(name)
	exist, err := translateNotFoundError(err)
	return secret, exist, err
}

func (c *clientWrapper) GetIngressClasses() ([]*netv1.IngressClass, error) {
	if c.clusterFactory == nil {
		return nil, errors.New("cluster factory not loaded")
	}

	var ics []*netv1.IngressClass
	ingressClasses, err := c.clusterFactory.Networking().V1().IngressClasses().Lister().List(labels.Everything())
	if err != nil {
		return nil, err
	}

	for _, ic := range ingressClasses {
		if ic.Spec.Controller == traefikDefaultIngressClassController {
			ics = append(ics, ic)
		}
	}

	return ics, nil
}

// lookupNamespace returns the lookup namespace key for the given namespace.
// When listening on all namespaces, it returns the client-go identifier ("")
// for all-namespaces. Otherwise, it returns the given namespace.
// The distinction is necessary because we index all informers on the special
// identifier iff all-namespaces are requested but receive specific namespace
// identifiers from the Kubernetes API, so we have to bridge this gap.
func (c *clientWrapper) lookupNamespace(ns string) string {
	if c.isNamespaceAll {
		return metav1.NamespaceAll
	}
	return ns
}

// translateNotFoundError will translate a "not found" error to a boolean return
// value which indicates if the resource exists and a nil error.
func translateNotFoundError(err error) (bool, error) {
	if kerror.IsNotFound(err) {
		return false, nil
	}
	return err == nil, err
}

// isWatchedNamespace checks to ensure that the namespace is being watched before we request
// it to ensure we don't panic by requesting an out-of-watch object.
func (c *clientWrapper) isWatchedNamespace(ns string) bool {
	if c.isNamespaceAll {
		return true
	}
	for _, watchedNamespace := range c.watchedNamespaces {
		if watchedNamespace == ns {
			return true
		}
	}
	return false
}

// filterIngressClassByName return a slice containing ingressclasses with the correct name.
func filterIngressClassByName(ingressClassName string, ics []*netv1.IngressClass) []*netv1.IngressClass {
	var ingressClasses []*netv1.IngressClass

	for _, ic := range ics {
		if ic.Name == ingressClassName {
			ingressClasses = append(ingressClasses, ic)
		}
	}

	return ingressClasses
}
