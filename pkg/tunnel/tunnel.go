package tunnel

import (
	"masterdnsvpn-go/internal/client"
	"masterdnsvpn-go/internal/config"
	"masterdnsvpn-go/internal/logger"
)

type (
	Client          = client.Client
	ClientConfig    = config.ClientConfig
	ResolverAddress = config.ResolverAddress
	Logger          = logger.Logger
)

func DiscardLogger() {
	logger.DiscardLogs()
}

func DefaultClientConfig() ClientConfig {
	return config.DefaultClientConfig()
}

func Bootstrap(cfg ClientConfig, path string) (*Client, error) {
	return client.BootstrapLoadedConfig(cfg, path)
}
