package config

type Option func(*Config)

func WithEnv(name string) Option {
	return func(c *Config) {
		c.env = name
	}
}
