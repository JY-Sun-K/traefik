<!--
CODE GENERATED AUTOMATICALLY
THIS FILE MUST NOT BE EDITED BY HAND
-->

`--accesslog`:  
Access log settings. (Default: ```false```)

`--accesslog.bufferingsize`:  
Number of access log lines to process in a buffered way. (Default: ```0```)

`--accesslog.fields.defaultmode`:  
Default mode for fields: keep | drop (Default: ```keep```)

`--accesslog.fields.headers.defaultmode`:  
Default mode for fields: keep | drop | redact (Default: ```drop```)

`--accesslog.fields.headers.names.<name>`:  
Override mode for headers

`--accesslog.fields.names.<name>`:  
Override mode for fields

`--accesslog.filepath`:  
Access log file path. Stdout is used when omitted or empty.

`--accesslog.filters.minduration`:  
Keep access logs when request took longer than the specified duration. (Default: ```0```)

`--accesslog.filters.retryattempts`:  
Keep access logs when at least one retry happened. (Default: ```false```)

`--accesslog.filters.statuscodes`:  
Keep access logs with status codes in the specified range.

`--accesslog.format`:  
Access log format: json | common (Default: ```common```)

`--api`:  
Enable api/dashboard. (Default: ```false```)

`--api.dashboard`:  
Activate dashboard. (Default: ```true```)

`--api.debug`:  
Enable additional endpoints for debugging and profiling. (Default: ```false```)

`--api.disabledashboardad`:  
Disable ad in the dashboard. (Default: ```false```)

`--api.insecure`:  
Activate API directly on the entryPoint named traefik. (Default: ```false```)

`--certificatesresolvers.<name>`:  
Certificates resolvers configuration. (Default: ```false```)

`--certificatesresolvers.<name>.acme.caserver`:  
CA server to use. (Default: ```https://acme-v02.api.letsencrypt.org/directory```)

`--certificatesresolvers.<name>.acme.certificatesduration`:  
Certificates' duration in hours. (Default: ```2160```)

`--certificatesresolvers.<name>.acme.dnschallenge`:  
Activate DNS-01 Challenge. (Default: ```false```)

`--certificatesresolvers.<name>.acme.dnschallenge.delaybeforecheck`:  
Assume DNS propagates after a delay in seconds rather than finding and querying nameservers. (Default: ```0```)

`--certificatesresolvers.<name>.acme.dnschallenge.disablepropagationcheck`:  
Disable the DNS propagation checks before notifying ACME that the DNS challenge is ready. [not recommended] (Default: ```false```)

`--certificatesresolvers.<name>.acme.dnschallenge.provider`:  
Use a DNS-01 based challenge provider rather than HTTPS.

`--certificatesresolvers.<name>.acme.dnschallenge.resolvers`:  
Use following DNS servers to resolve the FQDN authority.

`--certificatesresolvers.<name>.acme.eab.hmacencoded`:  
Base64 encoded HMAC key from External CA.

`--certificatesresolvers.<name>.acme.eab.kid`:  
Key identifier from External CA.

`--certificatesresolvers.<name>.acme.email`:  
Email address used for registration.

`--certificatesresolvers.<name>.acme.httpchallenge`:  
Activate HTTP-01 Challenge. (Default: ```false```)

`--certificatesresolvers.<name>.acme.httpchallenge.entrypoint`:  
HTTP challenge EntryPoint

`--certificatesresolvers.<name>.acme.keytype`:  
KeyType used for generating certificate private key. Allow value 'EC256', 'EC384', 'RSA2048', 'RSA4096', 'RSA8192'. (Default: ```RSA4096```)

`--certificatesresolvers.<name>.acme.preferredchain`:  
Preferred chain to use.

`--certificatesresolvers.<name>.acme.storage`:  
Storage to use. (Default: ```acme.json```)

`--certificatesresolvers.<name>.acme.tlschallenge`:  
Activate TLS-ALPN-01 Challenge. (Default: ```true```)

`--certificatesresolvers.<name>.tailscale`:  
Enables Tailscale certificate resolution. (Default: ```true```)

`--entrypoints.<name>`:  
Entry points definition. (Default: ```false```)

`--entrypoints.<name>.address`:  
Entry point address.

`--entrypoints.<name>.asdefault`:  
Adds this EntryPoint to the list of default EntryPoints to be used on routers that don't have any Entrypoint defined. (Default: ```false```)

`--entrypoints.<name>.forwardedheaders.insecure`:  
Trust all forwarded headers. (Default: ```false```)

`--entrypoints.<name>.forwardedheaders.trustedips`:  
Trust only forwarded headers from selected IPs.

`--entrypoints.<name>.http`:  
HTTP configuration.

`--entrypoints.<name>.http.encodequerysemicolons`:  
Defines whether request query semicolons should be URLEncoded. (Default: ```false```)

`--entrypoints.<name>.http.middlewares`:  
Default middlewares for the routers linked to the entry point.

`--entrypoints.<name>.http.redirections.entrypoint.permanent`:  
Applies a permanent redirection. (Default: ```true```)

`--entrypoints.<name>.http.redirections.entrypoint.priority`:  
Priority of the generated router. (Default: ```2147483646```)

`--entrypoints.<name>.http.redirections.entrypoint.scheme`:  
Scheme used for the redirection. (Default: ```https```)

`--entrypoints.<name>.http.redirections.entrypoint.to`:  
Targeted entry point of the redirection.

`--entrypoints.<name>.http.tls`:  
Default TLS configuration for the routers linked to the entry point. (Default: ```false```)

`--entrypoints.<name>.http.tls.certresolver`:  
Default certificate resolver for the routers linked to the entry point.

`--entrypoints.<name>.http.tls.domains`:  
Default TLS domains for the routers linked to the entry point.

`--entrypoints.<name>.http.tls.domains[n].main`:  
Default subject name.

`--entrypoints.<name>.http.tls.domains[n].sans`:  
Subject alternative names.

`--entrypoints.<name>.http.tls.options`:  
Default TLS options for the routers linked to the entry point.

`--entrypoints.<name>.http2.maxconcurrentstreams`:  
Specifies the number of concurrent streams per connection that each client is allowed to initiate. (Default: ```250```)

`--entrypoints.<name>.http3`:  
HTTP/3 configuration. (Default: ```false```)

`--entrypoints.<name>.http3.advertisedport`:  
UDP port to advertise, on which HTTP/3 is available. (Default: ```0```)

`--entrypoints.<name>.proxyprotocol`:  
Proxy-Protocol configuration. (Default: ```false```)

`--entrypoints.<name>.proxyprotocol.insecure`:  
Trust all. (Default: ```false```)

`--entrypoints.<name>.proxyprotocol.trustedips`:  
Trust only selected IPs.

`--entrypoints.<name>.transport.lifecycle.gracetimeout`:  
Duration to give active requests a chance to finish before Traefik stops. (Default: ```10```)

`--entrypoints.<name>.transport.lifecycle.requestacceptgracetimeout`:  
Duration to keep accepting requests before Traefik initiates the graceful shutdown procedure. (Default: ```0```)

`--entrypoints.<name>.transport.respondingtimeouts.idletimeout`:  
IdleTimeout is the maximum amount duration an idle (keep-alive) connection will remain idle before closing itself. If zero, no timeout is set. (Default: ```180```)

`--entrypoints.<name>.transport.respondingtimeouts.readtimeout`:  
ReadTimeout is the maximum duration for reading the entire request, including the body. If zero, no timeout is set. (Default: ```0```)

`--entrypoints.<name>.transport.respondingtimeouts.writetimeout`:  
WriteTimeout is the maximum duration before timing out writes of the response. If zero, no timeout is set. (Default: ```0```)

`--entrypoints.<name>.udp.timeout`:  
Timeout defines how long to wait on an idle session before releasing the related resources. (Default: ```3```)

`--experimental.http3`:  
Enable HTTP3. (Default: ```false```)

`--experimental.kubernetesgateway`:  
Allow the Kubernetes gateway api provider usage. (Default: ```false```)

`--experimental.localplugins.<name>`:  
Local plugins configuration. (Default: ```false```)

`--experimental.localplugins.<name>.modulename`:  
plugin's module name.

`--experimental.plugins.<name>.modulename`:  
plugin's module name.

`--experimental.plugins.<name>.version`:  
plugin's version.

`--global.checknewversion`:  
Periodically check if a new version has been released. (Default: ```true```)

`--global.sendanonymoususage`:  
Periodically send anonymous usage statistics. If the option is not specified, it will be enabled by default. (Default: ```false```)

`--hostresolver`:  
Enable CNAME Flattening. (Default: ```false```)

`--hostresolver.cnameflattening`:  
A flag to enable/disable CNAME flattening (Default: ```false```)

`--hostresolver.resolvconfig`:  
resolv.conf used for DNS resolving (Default: ```/etc/resolv.conf```)

`--hostresolver.resolvdepth`:  
The maximal depth of DNS recursive resolving (Default: ```5```)

`--log`:  
Traefik log settings. (Default: ```false```)

`--log.compress`:  
Determines if the rotated log files should be compressed using gzip. (Default: ```false```)

`--log.filepath`:  
Traefik log file path. Stdout is used when omitted or empty.

`--log.format`:  
Traefik log format: json | common (Default: ```common```)

`--log.level`:  
Log level set to traefik logs. (Default: ```ERROR```)

`--log.maxage`:  
Maximum number of days to retain old log files based on the timestamp encoded in their filename. (Default: ```0```)

`--log.maxbackups`:  
Maximum number of old log files to retain. (Default: ```0```)

`--log.maxsize`:  
Maximum size in megabytes of the log file before it gets rotated. (Default: ```0```)

`--log.nocolor`:  
When using the 'common' format, disables the colorized output. (Default: ```false```)

`--metrics.datadog`:  
Datadog metrics exporter type. (Default: ```false```)

`--metrics.datadog.addentrypointslabels`:  
Enable metrics on entry points. (Default: ```true```)

`--metrics.datadog.address`:  
Datadog's address. (Default: ```localhost:8125```)

`--metrics.datadog.addrouterslabels`:  
Enable metrics on routers. (Default: ```false```)

`--metrics.datadog.addserviceslabels`:  
Enable metrics on services. (Default: ```true```)

`--metrics.datadog.prefix`:  
Prefix to use for metrics collection. (Default: ```traefik```)

`--metrics.datadog.pushinterval`:  
Datadog push interval. (Default: ```10```)

`--metrics.influxdb2`:  
InfluxDB v2 metrics exporter type. (Default: ```false```)

`--metrics.influxdb2.addentrypointslabels`:  
Enable metrics on entry points. (Default: ```true```)

`--metrics.influxdb2.additionallabels.<name>`:  
Additional labels (influxdb tags) on all metrics

`--metrics.influxdb2.address`:  
InfluxDB v2 address. (Default: ```http://localhost:8086```)

`--metrics.influxdb2.addrouterslabels`:  
Enable metrics on routers. (Default: ```false```)

`--metrics.influxdb2.addserviceslabels`:  
Enable metrics on services. (Default: ```true```)

`--metrics.influxdb2.bucket`:  
InfluxDB v2 bucket ID.

`--metrics.influxdb2.org`:  
InfluxDB v2 org ID.

`--metrics.influxdb2.pushinterval`:  
InfluxDB v2 push interval. (Default: ```10```)

`--metrics.influxdb2.token`:  
InfluxDB v2 access token.

`--metrics.opentelemetry`:  
OpenTelemetry metrics exporter type. (Default: ```false```)

`--metrics.opentelemetry.addentrypointslabels`:  
Enable metrics on entry points. (Default: ```true```)

`--metrics.opentelemetry.address`:  
Address (host:port) of the collector endpoint. (Default: ```localhost:4318```)

`--metrics.opentelemetry.addrouterslabels`:  
Enable metrics on routers. (Default: ```false```)

`--metrics.opentelemetry.addserviceslabels`:  
Enable metrics on services. (Default: ```true```)

`--metrics.opentelemetry.explicitboundaries`:  
Boundaries for latency metrics. (Default: ```0.005000, 0.010000, 0.025000, 0.050000, 0.100000, 0.250000, 0.500000, 1.000000, 2.500000, 5.000000, 10.000000```)

`--metrics.opentelemetry.grpc`:  
gRPC specific configuration for the OpenTelemetry collector. (Default: ```true```)

`--metrics.opentelemetry.headers.<name>`:  
Headers sent with payload.

`--metrics.opentelemetry.insecure`:  
Disables client transport security for the exporter. (Default: ```false```)

`--metrics.opentelemetry.path`:  
Set the URL path of the collector endpoint.

`--metrics.opentelemetry.pushinterval`:  
Period between calls to collect a checkpoint. (Default: ```10```)

`--metrics.opentelemetry.tls.ca`:  
TLS CA

`--metrics.opentelemetry.tls.cert`:  
TLS cert

`--metrics.opentelemetry.tls.insecureskipverify`:  
TLS insecure skip verify (Default: ```false```)

`--metrics.opentelemetry.tls.key`:  
TLS key

`--metrics.prometheus`:  
Prometheus metrics exporter type. (Default: ```false```)

`--metrics.prometheus.addentrypointslabels`:  
Enable metrics on entry points. (Default: ```true```)

`--metrics.prometheus.addrouterslabels`:  
Enable metrics on routers. (Default: ```false```)

`--metrics.prometheus.addserviceslabels`:  
Enable metrics on services. (Default: ```true```)

`--metrics.prometheus.buckets`:  
Buckets for latency metrics. (Default: ```0.100000, 0.300000, 1.200000, 5.000000```)

`--metrics.prometheus.entrypoint`:  
EntryPoint (Default: ```traefik```)

`--metrics.prometheus.headerlabels.<name>`:  
Defines the extra labels for the requests_total metrics, and for each of them, the request header containing the value for this label.

`--metrics.prometheus.manualrouting`:  
Manual routing (Default: ```false```)

`--metrics.statsd`:  
StatsD metrics exporter type. (Default: ```false```)

`--metrics.statsd.addentrypointslabels`:  
Enable metrics on entry points. (Default: ```true```)

`--metrics.statsd.address`:  
StatsD address. (Default: ```localhost:8125```)

`--metrics.statsd.addrouterslabels`:  
Enable metrics on routers. (Default: ```false```)

`--metrics.statsd.addserviceslabels`:  
Enable metrics on services. (Default: ```true```)

`--metrics.statsd.prefix`:  
Prefix to use for metrics collection. (Default: ```traefik```)

`--metrics.statsd.pushinterval`:  
StatsD push interval. (Default: ```10```)

`--ping`:  
Enable ping. (Default: ```false```)

`--ping.entrypoint`:  
EntryPoint (Default: ```traefik```)

`--ping.manualrouting`:  
Manual routing (Default: ```false```)

`--ping.terminatingstatuscode`:  
Terminating status code (Default: ```503```)

`--providers.consul`:  
Enable Consul backend with default settings. (Default: ```false```)

`--providers.consul.endpoints`:  
KV store endpoints. (Default: ```127.0.0.1:8500```)

`--providers.consul.namespaces`:  
Sets the namespaces used to discover the configuration (Consul Enterprise only).

`--providers.consul.rootkey`:  
Root key used for KV store. (Default: ```traefik```)

`--providers.consul.tls.ca`:  
TLS CA

`--providers.consul.tls.cert`:  
TLS cert

`--providers.consul.tls.insecureskipverify`:  
TLS insecure skip verify (Default: ```false```)

`--providers.consul.tls.key`:  
TLS key

`--providers.consul.token`:  
Per-request ACL token.

`--providers.docker`:  
Enable Docker backend with default settings. (Default: ```false```)

`--providers.docker.allowemptyservices`:  
Disregards the Docker containers health checks with respect to the creation or removal of the corresponding services. (Default: ```false```)

`--providers.docker.constraints`:  
Constraints is an expression that Traefik matches against the container's labels to determine whether to create any route for that container.

`--providers.docker.defaultrule`:  
Default rule. (Default: ```Host(`{{ normalize .Name }}`)```)

`--providers.docker.endpoint`:  
Docker server endpoint. Can be a TCP or a Unix socket endpoint. (Default: ```unix:///var/run/docker.sock```)

`--providers.docker.exposedbydefault`:  
Expose containers by default. (Default: ```true```)

`--providers.docker.httpclienttimeout`:  
Client timeout for HTTP connections. (Default: ```0```)

`--providers.docker.network`:  
Default Docker network used.

`--providers.docker.tls.ca`:  
TLS CA

`--providers.docker.tls.cert`:  
TLS cert

`--providers.docker.tls.insecureskipverify`:  
TLS insecure skip verify (Default: ```false```)

`--providers.docker.tls.key`:  
TLS key

`--providers.docker.usebindportip`:  
Use the ip address from the bound port, rather than from the inner network. (Default: ```false```)

`--providers.docker.watch`:  
Watch Docker events. (Default: ```true```)

`--providers.ecs`:  
Enable AWS ECS backend with default settings. (Default: ```false```)

`--providers.ecs.accesskeyid`:  
AWS credentials access key ID to use for making requests.

`--providers.ecs.autodiscoverclusters`:  
Auto discover cluster. (Default: ```false```)

`--providers.ecs.clusters`:  
ECS Cluster names. (Default: ```default```)

`--providers.ecs.constraints`:  
Constraints is an expression that Traefik matches against the container's labels to determine whether to create any route for that container.

`--providers.ecs.defaultrule`:  
Default rule. (Default: ```Host(`{{ normalize .Name }}`)```)

`--providers.ecs.ecsanywhere`:  
Enable ECS Anywhere support. (Default: ```false```)

`--providers.ecs.exposedbydefault`:  
Expose services by default. (Default: ```true```)

`--providers.ecs.healthytasksonly`:  
Determines whether to discover only healthy tasks. (Default: ```false```)

`--providers.ecs.refreshseconds`:  
Polling interval (in seconds). (Default: ```15```)

`--providers.ecs.region`:  
AWS region to use for requests.

`--providers.ecs.secretaccesskey`:  
AWS credentials access key to use for making requests.

`--providers.etcd`:  
Enable Etcd backend with default settings. (Default: ```false```)

`--providers.etcd.endpoints`:  
KV store endpoints. (Default: ```127.0.0.1:2379```)

`--providers.etcd.password`:  
Password for authentication.

`--providers.etcd.rootkey`:  
Root key used for KV store. (Default: ```traefik```)

`--providers.etcd.tls.ca`:  
TLS CA

`--providers.etcd.tls.cert`:  
TLS cert

`--providers.etcd.tls.insecureskipverify`:  
TLS insecure skip verify (Default: ```false```)

`--providers.etcd.tls.key`:  
TLS key

`--providers.etcd.username`:  
Username for authentication.

`--providers.file.debugloggeneratedtemplate`:  
Enable debug logging of generated configuration template. (Default: ```false```)

`--providers.file.directory`:  
Load dynamic configuration from one or more .yml or .toml files in a directory.

`--providers.file.filename`:  
Load dynamic configuration from a file.

`--providers.file.watch`:  
Watch provider. (Default: ```true```)

`--providers.http`:  
Enable HTTP backend with default settings. (Default: ```false```)

`--providers.http.endpoint`:  
Load configuration from this endpoint.

`--providers.http.headers.<name>`:  
Define custom headers to be sent to the endpoint.

`--providers.http.pollinterval`:  
Polling interval for endpoint. (Default: ```5```)

`--providers.http.polltimeout`:  
Polling timeout for endpoint. (Default: ```5```)

`--providers.http.tls.ca`:  
TLS CA

`--providers.http.tls.cert`:  
TLS cert

`--providers.http.tls.insecureskipverify`:  
TLS insecure skip verify (Default: ```false```)

`--providers.http.tls.key`:  
TLS key

`--providers.kubernetescrd`:  
Enable Kubernetes backend with default settings. (Default: ```false```)

`--providers.kubernetescrd.allowcrossnamespace`:  
Allow cross namespace resource reference. (Default: ```false```)

`--providers.kubernetescrd.allowemptyservices`:  
Allow the creation of services without endpoints. (Default: ```false```)

`--providers.kubernetescrd.allowexternalnameservices`:  
Allow ExternalName services. (Default: ```false```)

`--providers.kubernetescrd.certauthfilepath`:  
Kubernetes certificate authority file path (not needed for in-cluster client).

`--providers.kubernetescrd.endpoint`:  
Kubernetes server endpoint (required for external cluster client).

`--providers.kubernetescrd.ingressclass`:  
Value of kubernetes.io/ingress.class annotation to watch for.

`--providers.kubernetescrd.labelselector`:  
Kubernetes label selector to use.

`--providers.kubernetescrd.namespaces`:  
Kubernetes namespaces.

`--providers.kubernetescrd.throttleduration`:  
Ingress refresh throttle duration (Default: ```0```)

`--providers.kubernetescrd.token`:  
Kubernetes bearer token (not needed for in-cluster client).

`--providers.kubernetesgateway`:  
Enable Kubernetes gateway api provider with default settings. (Default: ```false```)

`--providers.kubernetesgateway.certauthfilepath`:  
Kubernetes certificate authority file path (not needed for in-cluster client).

`--providers.kubernetesgateway.endpoint`:  
Kubernetes server endpoint (required for external cluster client).

`--providers.kubernetesgateway.labelselector`:  
Kubernetes label selector to select specific GatewayClasses.

`--providers.kubernetesgateway.namespaces`:  
Kubernetes namespaces.

`--providers.kubernetesgateway.throttleduration`:  
Kubernetes refresh throttle duration (Default: ```0```)

`--providers.kubernetesgateway.token`:  
Kubernetes bearer token (not needed for in-cluster client).

`--providers.kubernetesingress`:  
Enable Kubernetes backend with default settings. (Default: ```false```)

`--providers.kubernetesingress.allowemptyservices`:  
Allow creation of services without endpoints. (Default: ```false```)

`--providers.kubernetesingress.allowexternalnameservices`:  
Allow ExternalName services. (Default: ```false```)

`--providers.kubernetesingress.certauthfilepath`:  
Kubernetes certificate authority file path (not needed for in-cluster client).

`--providers.kubernetesingress.disableingressclasslookup`:  
Disables the lookup of IngressClasses. (Default: ```false```)

`--providers.kubernetesingress.endpoint`:  
Kubernetes server endpoint (required for external cluster client).

`--providers.kubernetesingress.ingressclass`:  
Value of kubernetes.io/ingress.class annotation or IngressClass name to watch for.

`--providers.kubernetesingress.ingressendpoint.hostname`:  
Hostname used for Kubernetes Ingress endpoints.

`--providers.kubernetesingress.ingressendpoint.ip`:  
IP used for Kubernetes Ingress endpoints.

`--providers.kubernetesingress.ingressendpoint.publishedservice`:  
Published Kubernetes Service to copy status from.

`--providers.kubernetesingress.labelselector`:  
Kubernetes Ingress label selector to use.

`--providers.kubernetesingress.namespaces`:  
Kubernetes namespaces.

`--providers.kubernetesingress.throttleduration`:  
Ingress refresh throttle duration (Default: ```0```)

`--providers.kubernetesingress.token`:  
Kubernetes bearer token (not needed for in-cluster client).

`--providers.nomad`:  
Enable Nomad backend with default settings. (Default: ```false```)

`--providers.nomad.constraints`:  
Constraints is an expression that Traefik matches against the Nomad service's tags to determine whether to create route(s) for that service.

`--providers.nomad.defaultrule`:  
Default rule. (Default: ```Host(`{{ normalize .Name }}`)```)

`--providers.nomad.endpoint.address`:  
The address of the Nomad server, including scheme and port. (Default: ```http://127.0.0.1:4646```)

`--providers.nomad.endpoint.endpointwaittime`:  
WaitTime limits how long a Watch will block. If not provided, the agent default values will be used (Default: ```0```)

`--providers.nomad.endpoint.region`:  
Nomad region to use. If not provided, the local agent region is used.

`--providers.nomad.endpoint.tls.ca`:  
TLS CA

`--providers.nomad.endpoint.tls.cert`:  
TLS cert

`--providers.nomad.endpoint.tls.insecureskipverify`:  
TLS insecure skip verify (Default: ```false```)

`--providers.nomad.endpoint.tls.key`:  
TLS key

`--providers.nomad.endpoint.token`:  
Token is used to provide a per-request ACL token.

`--providers.nomad.exposedbydefault`:  
Expose Nomad services by default. (Default: ```true```)

`--providers.nomad.namespaces`:  
Sets the Nomad namespaces used to discover services.

`--providers.nomad.prefix`:  
Prefix for nomad service tags. (Default: ```traefik```)

`--providers.nomad.refreshinterval`:  
Interval for polling Nomad API. (Default: ```15```)

`--providers.nomad.stale`:  
Use stale consistency for catalog reads. (Default: ```false```)

`--providers.plugin.<name>`:  
Plugins configuration.

`--providers.providersthrottleduration`:  
Backends throttle duration: minimum duration between 2 events from providers before applying a new configuration. It avoids unnecessary reloads if multiples events are sent in a short amount of time. (Default: ```2```)

`--providers.redis`:  
Enable Redis backend with default settings. (Default: ```false```)

`--providers.redis.db`:  
Database to be selected after connecting to the server. (Default: ```0```)

`--providers.redis.endpoints`:  
KV store endpoints. (Default: ```127.0.0.1:6379```)

`--providers.redis.password`:  
Password for authentication.

`--providers.redis.rootkey`:  
Root key used for KV store. (Default: ```traefik```)

`--providers.redis.tls.ca`:  
TLS CA

`--providers.redis.tls.cert`:  
TLS cert

`--providers.redis.tls.insecureskipverify`:  
TLS insecure skip verify (Default: ```false```)

`--providers.redis.tls.key`:  
TLS key

`--providers.redis.username`:  
Username for authentication.

`--providers.rest`:  
Enable Rest backend with default settings. (Default: ```false```)

`--providers.rest.insecure`:  
Activate REST Provider directly on the entryPoint named traefik. (Default: ```false```)

`--providers.swarm`:  
Enable Docker Swarm backend with default settings. (Default: ```false```)

`--providers.swarm.allowemptyservices`:  
Disregards the Docker containers health checks with respect to the creation or removal of the corresponding services. (Default: ```false```)

`--providers.swarm.constraints`:  
Constraints is an expression that Traefik matches against the container's labels to determine whether to create any route for that container.

`--providers.swarm.defaultrule`:  
Default rule. (Default: ```Host(`{{ normalize .Name }}`)```)

`--providers.swarm.endpoint`:  
Docker server endpoint. Can be a TCP or a Unix socket endpoint. (Default: ```unix:///var/run/docker.sock```)

`--providers.swarm.exposedbydefault`:  
Expose containers by default. (Default: ```true```)

`--providers.swarm.httpclienttimeout`:  
Client timeout for HTTP connections. (Default: ```0```)

`--providers.swarm.network`:  
Default Docker network used.

`--providers.swarm.refreshseconds`:  
Polling interval for swarm mode. (Default: ```15```)

`--providers.swarm.tls.ca`:  
TLS CA

`--providers.swarm.tls.cert`:  
TLS cert

`--providers.swarm.tls.insecureskipverify`:  
TLS insecure skip verify (Default: ```false```)

`--providers.swarm.tls.key`:  
TLS key

`--providers.swarm.usebindportip`:  
Use the ip address from the bound port, rather than from the inner network. (Default: ```false```)

`--providers.swarm.watch`:  
Watch Docker events. (Default: ```true```)

`--providers.zookeeper`:  
Enable ZooKeeper backend with default settings. (Default: ```false```)

`--providers.zookeeper.endpoints`:  
KV store endpoints. (Default: ```127.0.0.1:2181```)

`--providers.zookeeper.password`:  
Password for authentication.

`--providers.zookeeper.rootkey`:  
Root key used for KV store. (Default: ```traefik```)

`--providers.zookeeper.username`:  
Username for authentication.

`--serverstransport.forwardingtimeouts.dialtimeout`:  
The amount of time to wait until a connection to a backend server can be established. If zero, no timeout exists. (Default: ```30```)

`--serverstransport.forwardingtimeouts.idleconntimeout`:  
The maximum period for which an idle HTTP keep-alive connection will remain open before closing itself (Default: ```90```)

`--serverstransport.forwardingtimeouts.responseheadertimeout`:  
The amount of time to wait for a server's response headers after fully writing the request (including its body, if any). If zero, no timeout exists. (Default: ```0```)

`--serverstransport.insecureskipverify`:  
Disable SSL certificate verification. (Default: ```false```)

`--serverstransport.maxidleconnsperhost`:  
If non-zero, controls the maximum idle (keep-alive) to keep per-host. If zero, DefaultMaxIdleConnsPerHost is used (Default: ```200```)

`--serverstransport.rootcas`:  
Add cert file for self-signed certificate.

`--serverstransport.spiffe`:  
Defines the SPIFFE configuration. (Default: ```false```)

`--serverstransport.spiffe.ids`:  
Defines the allowed SPIFFE IDs (takes precedence over the SPIFFE TrustDomain).

`--serverstransport.spiffe.trustdomain`:  
Defines the allowed SPIFFE trust domain.

`--spiffe.workloadapiaddr`:  
Defines the workload API address.

`--tcpserverstransport.dialkeepalive`:  
Defines the interval between keep-alive probes for an active network connection. If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled (Default: ```15```)

`--tcpserverstransport.dialtimeout`:  
Defines the amount of time to wait until a connection to a backend server can be established. If zero, no timeout exists. (Default: ```30```)

`--tcpserverstransport.terminationdelay`:  
Defines the delay to wait before fully terminating the connection, after one connected peer has closed its writing capability. (Default: ```0```)

`--tcpserverstransport.tls`:  
Defines the TLS configuration. (Default: ```false```)

`--tcpserverstransport.tls.insecureskipverify`:  
Disables SSL certificate verification. (Default: ```false```)

`--tcpserverstransport.tls.rootcas`:  
Defines a list of CA secret used to validate self-signed certificate

`--tcpserverstransport.tls.spiffe`:  
Defines the SPIFFE TLS configuration. (Default: ```false```)

`--tcpserverstransport.tls.spiffe.ids`:  
Defines the allowed SPIFFE IDs (takes precedence over the SPIFFE TrustDomain).

`--tcpserverstransport.tls.spiffe.trustdomain`:  
Defines the allowed SPIFFE trust domain.

`--tracing`:  
OpenTracing configuration. (Default: ```false```)

`--tracing.datadog`:  
Settings for Datadog. (Default: ```false```)

`--tracing.datadog.bagageprefixheadername`:  
Sets the header name prefix used to store baggage items in a map.

`--tracing.datadog.debug`:  
Enables Datadog debug. (Default: ```false```)

`--tracing.datadog.globaltags.<name>`:  
Sets a list of key:value tags on all spans.

`--tracing.datadog.localagenthostport`:  
Sets the Datadog Agent host:port. (Default: ```localhost:8126```)

`--tracing.datadog.localagentsocket`:  
Sets the socket for the Datadog Agent.

`--tracing.datadog.parentidheadername`:  
Sets the header name used to store the parent ID.

`--tracing.datadog.prioritysampling`:  
Enables priority sampling. When using distributed tracing, this option must be enabled in order to get all the parts of a distributed trace sampled. (Default: ```false```)

`--tracing.datadog.samplingpriorityheadername`:  
Sets the header name used to store the sampling priority.

`--tracing.datadog.traceidheadername`:  
Sets the header name used to store the trace ID.

`--tracing.elastic`:  
Settings for Elastic. (Default: ```false```)

`--tracing.elastic.secrettoken`:  
Sets the token used to connect to Elastic APM Server.

`--tracing.elastic.serverurl`:  
Sets the URL of the Elastic APM server.

`--tracing.elastic.serviceenvironment`:  
Sets the name of the environment Traefik is deployed in, e.g. 'production' or 'staging'.

`--tracing.haystack`:  
Settings for Haystack. (Default: ```false```)

`--tracing.haystack.baggageprefixheadername`:  
Sets the header name prefix used to store baggage items in a map.

`--tracing.haystack.globaltag`:  
Sets a key:value tag on all spans.

`--tracing.haystack.localagenthost`:  
Sets the Haystack Agent host. (Default: ```127.0.0.1```)

`--tracing.haystack.localagentport`:  
Sets the Haystack Agent port. (Default: ```35000```)

`--tracing.haystack.parentidheadername`:  
Sets the header name used to store the parent ID.

`--tracing.haystack.spanidheadername`:  
Sets the header name used to store the span ID.

`--tracing.haystack.traceidheadername`:  
Sets the header name used to store the trace ID.

`--tracing.instana`:  
Settings for Instana. (Default: ```false```)

`--tracing.instana.enableautoprofile`:  
Enables automatic profiling for the Traefik process. (Default: ```false```)

`--tracing.instana.localagenthost`:  
Sets the Instana Agent host.

`--tracing.instana.localagentport`:  
Sets the Instana Agent port. (Default: ```42699```)

`--tracing.instana.loglevel`:  
Sets the log level for the Instana tracer. ('error','warn','info','debug') (Default: ```info```)

`--tracing.jaeger`:  
Settings for Jaeger. (Default: ```false```)

`--tracing.jaeger.collector.endpoint`:  
Instructs reporter to send spans to jaeger-collector at this URL.

`--tracing.jaeger.collector.password`:  
Password for basic http authentication when sending spans to jaeger-collector.

`--tracing.jaeger.collector.user`:  
User for basic http authentication when sending spans to jaeger-collector.

`--tracing.jaeger.disableattemptreconnecting`:  
Disables the periodic re-resolution of the agent's hostname and reconnection if there was a change. (Default: ```true```)

`--tracing.jaeger.gen128bit`:  
Generates 128 bits span IDs. (Default: ```false```)

`--tracing.jaeger.localagenthostport`:  
Sets the Jaeger Agent host:port. (Default: ```127.0.0.1:6831```)

`--tracing.jaeger.propagation`:  
Sets the propagation format (jaeger/b3). (Default: ```jaeger```)

`--tracing.jaeger.samplingparam`:  
Sets the sampling parameter. (Default: ```1.000000```)

`--tracing.jaeger.samplingserverurl`:  
Sets the sampling server URL. (Default: ```http://localhost:5778/sampling```)

`--tracing.jaeger.samplingtype`:  
Sets the sampling type. (Default: ```const```)

`--tracing.jaeger.tracecontextheadername`:  
Sets the header name used to store the trace ID. (Default: ```uber-trace-id```)

`--tracing.opentelemetry`:  
Settings for OpenTelemetry. (Default: ```false```)

`--tracing.opentelemetry.address`:  
Sets the address (host:port) of the collector endpoint. (Default: ```localhost:4318```)

`--tracing.opentelemetry.grpc`:  
gRPC specific configuration for the OpenTelemetry collector. (Default: ```true```)

`--tracing.opentelemetry.headers.<name>`:  
Defines additional headers to be sent with the payloads.

`--tracing.opentelemetry.insecure`:  
Disables client transport security for the exporter. (Default: ```false```)

`--tracing.opentelemetry.path`:  
Sets the URL path of the collector endpoint.

`--tracing.opentelemetry.tls.ca`:  
TLS CA

`--tracing.opentelemetry.tls.cert`:  
TLS cert

`--tracing.opentelemetry.tls.insecureskipverify`:  
TLS insecure skip verify (Default: ```false```)

`--tracing.opentelemetry.tls.key`:  
TLS key

`--tracing.servicename`:  
Set the name for this service. (Default: ```traefik```)

`--tracing.spannamelimit`:  
Set the maximum character limit for Span names (default 0 = no limit). (Default: ```0```)

`--tracing.zipkin`:  
Settings for Zipkin. (Default: ```false```)

`--tracing.zipkin.httpendpoint`:  
Sets the HTTP Endpoint to report traces to. (Default: ```http://localhost:9411/api/v2/spans```)

`--tracing.zipkin.id128bit`:  
Uses 128 bits root span IDs. (Default: ```true```)

`--tracing.zipkin.samespan`:  
Uses SameSpan RPC style traces. (Default: ```false```)

`--tracing.zipkin.samplerate`:  
Sets the rate between 0.0 and 1.0 of requests to trace. (Default: ```1.000000```)
