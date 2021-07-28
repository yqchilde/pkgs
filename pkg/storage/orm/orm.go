package orm

import "time"

type Config struct {
	Addr            string        `mapstructure:"addr"`
	Database        string        `mapstructure:"database"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	ShowLog         bool          `mapstructure:"show-log"`
	AutoMigrate     bool          `mapstructure:"auto-migrate"`
	MaxIdleConn     int           `mapstructure:"max-idle-conn"`
	MaxOpenConn     int           `mapstructure:"max-open-conn"`
	ConnMaxLifeTime time.Duration `mapstructure:"conn-max-life-time"`
}