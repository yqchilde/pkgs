package config

import "github.com/yqchilde/gint/pkg/logger"

// Config for application config
type Config struct {
	Server ServerConfig
	Http APIConfig
	Grpc APIConfig

	// component config
	Logger logger.Config
}
