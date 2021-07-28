package config

import (
	"github.com/yqchilde/gint/pkg/logger"
	"github.com/yqchilde/gint/pkg/redis"
	"github.com/yqchilde/gint/pkg/storage/orm"
)

// Config for application config
type Config struct {
	Server ServerConfig
	Http   APIConfig
	Grpc   APIConfig

	// component config
	Logger logger.Config
	MySQL  orm.Config
	Redis  redis.Config
}
