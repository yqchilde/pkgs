package orm

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/yqchilde/gin-skeleton/pkg/log"
)

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

// NewMySQL generate mysql orm instance
func NewMySQL(c *Config) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		c.Username,
		c.Password,
		c.Addr,
		c.Database,
		true,
		"Local")

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panicf("open mysql failed. database name: %s, err: %+v", c.Database, err)
	}
	// set for db connection
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(c.ConnMaxLifeTime)

	db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), gormConfig(c))
	if err != nil {
		log.Panicf("database connection failed. database name: %s, err: %+v", c.Database, err)
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	return db
}

func gormConfig(c *Config) *gorm.Config {
	if !c.ShowLog {
		return &gorm.Config{}
	}

	return &gorm.Config{
		Logger:                                   gormLogger.Default.LogMode(gormLogger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
