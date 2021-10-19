package conf

import (
	"time"

	"github.com/yqchilde/gin-skeleton/pkg/log"
	"github.com/yqchilde/gin-skeleton/pkg/redis"
	"github.com/yqchilde/gin-skeleton/pkg/storage/orm"
)

// Config for application conf
type Config struct {
	App  AppConfig
	HTTP ServerConfig
	Grpc ServerConfig

	// component conf
	Logger log.Config
	MySQL  orm.Config
	Redis  redis.Config
}

// AppConfig for application api conf
type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Mode    string `mapstructure:"mode"`
	SSL     bool   `mapstructure:"ssl"`
	Debug   bool   `mapstructure:"debug"`

	// AuthConfig
	JwtSecret       string `mapstructure:"jwt-secret"`
	JwtExpireTime   int64  `mapstructure:"jwt-expire-time"`
	CasbinModelPath string `mapstructure:"casbin-model-path"`
}

type ServerConfig struct {
	Network      string        `mapstructure:"network"`
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout"`
}
