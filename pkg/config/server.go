package config

type ServerConfig struct {
	Name  string `mapstructure:"name"`
	Mode  string `mapstructure:"mode"`
	URL   string `mapstructure:"url"`
	SSL   bool   `mapstructure:"ssl"`
	Debug bool   `mapstructure:"debug"`

	AuthConfig
}

type AuthConfig struct {
	JwtSecret       string `mapstructure:"jwt-secret"`
	JwtExpireTime   int64  `mapstructure:"jwt-expire-time"`
	CasbinModelPath string `mapstructure:"casbin-model-path"`
}
