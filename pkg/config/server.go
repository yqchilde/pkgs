package config

// ServerConfig for application server config
type ServerConfig struct {
	Name  string `mapstructure:"name"`
	Mode  string `mapstructure:"mode"`
	SSL   bool   `mapstructure:"ssl"`
	Debug bool   `mapstructure:"debug"`

	// AuthConfig
	JwtSecret       string `mapstructure:"jwt-secret"`
	JwtExpireTime   int64  `mapstructure:"jwt-expire-time"`
	CasbinModelPath string `mapstructure:"casbin-model-path"`
}
