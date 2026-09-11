package tunnel

import (
	"masterdnsvpn-go/internal/client"
	"masterdnsvpn-go/internal/config"
)

type (
	Client          = client.Client
	ClientConfig    = config.ClientConfig
	ResolverAddress = config.ResolverAddress
)

func DefaultClientConfig() ClientConfig {
	return config.DefaultClientConfig()
}

func Bootstrap(cfg ClientConfig, path string) (*Client, error) {
	return client.BootstrapLoadedConfig(cfg, path)
}
