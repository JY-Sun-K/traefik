package cmd

import (
	"time"
	"traefik/v3/pkg/provider/http"

	ptypes "github.com/traefik/paerser/types"
	"traefik/v3/pkg/config/static"
)

// TraefikCmdConfiguration wraps the static configuration and extra parameters.
type TraefikCmdConfiguration struct {
	static.Configuration `export:"true"`
	// ConfigFile is the path to the configuration file.
	ConfigFile string `description:"Configuration file to use. If specified all other flags are ignored." export:"true"`
}

// NewTraefikConfiguration creates a TraefikCmdConfiguration with default values.
func NewTraefikConfiguration() *TraefikCmdConfiguration {
	httpProvider := &http.Provider{
		Endpoint:     "http://127.0.0.1:9000",
		PollInterval: 0,
		PollTimeout:  0,
		Headers:      nil,
		TLS:          nil,
	}

	httpProvider.SetDefaults()
	httpProvider.Init()
	end := make(static.EntryPoints)
	eq := &static.EntryPoint{
		Address: ":8000",
	}
	eq.SetDefaults()
	end["web"] = eq
	return &TraefikCmdConfiguration{
		Configuration: static.Configuration{
			Global: &static.Global{
				CheckNewVersion: true,
			},
			EntryPoints: end,
			Providers: &static.Providers{
				ProvidersThrottleDuration: ptypes.Duration(2 * time.Second),
				HTTP:                      httpProvider,
			},
			ServersTransport: &static.ServersTransport{
				MaxIdleConnsPerHost: 200,
			},
			TCPServersTransport: &static.TCPServersTransport{
				DialTimeout:   ptypes.Duration(30 * time.Second),
				DialKeepAlive: ptypes.Duration(15 * time.Second),
			},
			API: &static.API{
				Insecure:  true,
				Dashboard: true,
				Debug:     true,
			},
		},
		ConfigFile: "",
	}
}
