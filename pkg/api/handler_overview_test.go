package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"traefik/v3/pkg/config/dynamic"
	"traefik/v3/pkg/config/runtime"
	"traefik/v3/pkg/config/static"
	"traefik/v3/pkg/provider/docker"
	"traefik/v3/pkg/provider/file"
	"traefik/v3/pkg/provider/kubernetes/crd"
	"traefik/v3/pkg/provider/kubernetes/ingress"
	"traefik/v3/pkg/provider/rest"
	"traefik/v3/pkg/tracing/jaeger"
	"traefik/v3/pkg/types"
)

func TestHandler_Overview(t *testing.T) {
	type expected struct {
		statusCode int
		jsonFile   string
	}

	testCases := []struct {
		desc       string
		path       string
		confStatic static.Configuration
		confDyn    runtime.Configuration
		expected   expected
	}{
		{
			desc:       "without data in the dynamic configuration",
			path:       "/api/overview",
			confStatic: static.Configuration{API: &static.API{}, Global: &static.Global{}},
			confDyn:    runtime.Configuration{},
			expected: expected{
				statusCode: http.StatusOK,
				jsonFile:   "testdata/overview-empty.json",
			},
		},
		{
			desc:       "with data in the dynamic configuration",
			path:       "/api/overview",
			confStatic: static.Configuration{API: &static.API{}, Global: &static.Global{}},
			confDyn: runtime.Configuration{
				Services: map[string]*runtime.ServiceInfo{
					"foo-service@myprovider": {
						Service: &dynamic.Service{
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Servers: []dynamic.Server{{URL: "http://127.0.0.1"}},
							},
						},
						Status: runtime.StatusEnabled,
					},
					"bar-service@myprovider": {
						Service: &dynamic.Service{
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Servers: []dynamic.Server{{URL: "http://127.0.0.1"}},
							},
						},
						Status: runtime.StatusWarning,
					},
					"fii-service@myprovider": {
						Service: &dynamic.Service{
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Servers: []dynamic.Server{{URL: "http://127.0.0.1"}},
							},
						},
						Status: runtime.StatusDisabled,
					},
				},
				Middlewares: map[string]*runtime.MiddlewareInfo{
					"auth@myprovider": {
						Middleware: &dynamic.Middleware{
							BasicAuth: &dynamic.BasicAuth{
								Users: []string{"admin:admin"},
							},
						},
						Status: runtime.StatusEnabled,
					},
					"addPrefixTest@myprovider": {
						Middleware: &dynamic.Middleware{
							AddPrefix: &dynamic.AddPrefix{
								Prefix: "/titi",
							},
						},
					},
					"addPrefixTest@anotherprovider": {
						Middleware: &dynamic.Middleware{
							AddPrefix: &dynamic.AddPrefix{
								Prefix: "/toto",
							},
						},
						Err:    []string{"error"},
						Status: runtime.StatusDisabled,
					},
				},
				Routers: map[string]*runtime.RouterInfo{
					"bar@myprovider": {
						Router: &dynamic.Router{
							EntryPoints: []string{"web"},
							Service:     "foo-service@myprovider",
							Rule:        "Host(`foo.bar`)",
							Middlewares: []string{"auth", "addPrefixTest@anotherprovider"},
						},
						Status: runtime.StatusEnabled,
					},
					"test@myprovider": {
						Router: &dynamic.Router{
							EntryPoints: []string{"web"},
							Service:     "foo-service@myprovider",
							Rule:        "Host(`foo.bar.other`)",
							Middlewares: []string{"addPrefixTest", "auth"},
						},
						Status: runtime.StatusWarning,
					},
					"foo@myprovider": {
						Router: &dynamic.Router{
							EntryPoints: []string{"web"},
							Service:     "foo-service@myprovider",
							Rule:        "Host(`foo.bar.other`)",
							Middlewares: []string{"addPrefixTest", "auth"},
						},
						Status: runtime.StatusDisabled,
					},
				},
				TCPServices: map[string]*runtime.TCPServiceInfo{
					"tcpfoo-service@myprovider": {
						TCPService: &dynamic.TCPService{
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "127.0.0.1",
									},
								},
							},
						},
						Status: runtime.StatusEnabled,
					},
					"tcpbar-service@myprovider": {
						TCPService: &dynamic.TCPService{
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "127.0.0.2",
									},
								},
							},
						},
						Status: runtime.StatusWarning,
					},
					"tcpfii-service@myprovider": {
						TCPService: &dynamic.TCPService{
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "127.0.0.2",
									},
								},
							},
						},
						Status: runtime.StatusDisabled,
					},
				},
				TCPMiddlewares: map[string]*runtime.TCPMiddlewareInfo{
					"ipallowlist1@myprovider": {
						TCPMiddleware: &dynamic.TCPMiddleware{
							IPAllowList: &dynamic.TCPIPAllowList{
								SourceRange: []string{"127.0.0.1/32"},
							},
						},
						Status: runtime.StatusEnabled,
					},
					"ipallowlist2@myprovider": {
						TCPMiddleware: &dynamic.TCPMiddleware{
							IPAllowList: &dynamic.TCPIPAllowList{
								SourceRange: []string{"127.0.0.1/32"},
							},
						},
					},
					"ipallowlist3@myprovider": {
						TCPMiddleware: &dynamic.TCPMiddleware{
							IPAllowList: &dynamic.TCPIPAllowList{
								SourceRange: []string{"127.0.0.1/32"},
							},
						},
						Status: runtime.StatusDisabled,
					},
				},
				TCPRouters: map[string]*runtime.TCPRouterInfo{
					"tcpbar@myprovider": {
						TCPRouter: &dynamic.TCPRouter{
							EntryPoints: []string{"web"},
							Service:     "tcpfoo-service@myprovider",
							Rule:        "HostSNI(`foo.bar`)",
						},
						Status: runtime.StatusEnabled,
					},
					"tcptest@myprovider": {
						TCPRouter: &dynamic.TCPRouter{
							EntryPoints: []string{"web"},
							Service:     "tcpfoo-service@myprovider",
							Rule:        "HostSNI(`foo.bar.other`)",
						},
						Status: runtime.StatusWarning,
					},
					"tcpfoo@myprovider": {
						TCPRouter: &dynamic.TCPRouter{
							EntryPoints: []string{"web"},
							Service:     "tcpfoo-service@myprovider",
							Rule:        "HostSNI(`foo.bar.other`)",
						},
						Status: runtime.StatusDisabled,
					},
				},
			},
			expected: expected{
				statusCode: http.StatusOK,
				jsonFile:   "testdata/overview-dynamic.json",
			},
		},
		{
			desc: "with providers",
			path: "/api/overview",
			confStatic: static.Configuration{
				Global: &static.Global{},
				API:    &static.API{},
				Providers: &static.Providers{
					Docker:            &docker.Provider{},
					Swarm:             &docker.SwarmProvider{},
					File:              &file.Provider{},
					KubernetesIngress: &ingress.Provider{},
					KubernetesCRD:     &crd.Provider{},
					Rest:              &rest.Provider{},
					Plugin: map[string]static.PluginConf{
						"test": map[string]interface{}{},
					},
				},
			},
			confDyn: runtime.Configuration{},
			expected: expected{
				statusCode: http.StatusOK,
				jsonFile:   "testdata/overview-providers.json",
			},
		},
		{
			desc: "with features",
			path: "/api/overview",
			confStatic: static.Configuration{
				Global: &static.Global{},
				API:    &static.API{},
				Metrics: &types.Metrics{
					Prometheus: &types.Prometheus{},
				},
				Tracing: &static.Tracing{
					Jaeger: &jaeger.Config{},
				},
			},
			confDyn: runtime.Configuration{},
			expected: expected{
				statusCode: http.StatusOK,
				jsonFile:   "testdata/overview-features.json",
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			handler := New(test.confStatic, &test.confDyn)
			server := httptest.NewServer(handler.createRouter())

			resp, err := http.DefaultClient.Get(server.URL + test.path)
			require.NoError(t, err)

			require.Equal(t, test.expected.statusCode, resp.StatusCode)

			if test.expected.jsonFile == "" {
				return
			}

			assert.Equal(t, resp.Header.Get("Content-Type"), "application/json")
			contents, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			err = resp.Body.Close()
			require.NoError(t, err)

			if *updateExpected {
				var results interface{}
				err := json.Unmarshal(contents, &results)
				require.NoError(t, err)

				newJSON, err := json.MarshalIndent(results, "", "\t")
				require.NoError(t, err)

				err = os.WriteFile(test.expected.jsonFile, newJSON, 0o644)
				require.NoError(t, err)
			}

			data, err := os.ReadFile(test.expected.jsonFile)
			require.NoError(t, err)
			assert.JSONEq(t, string(data), string(contents))
		})
	}
}
