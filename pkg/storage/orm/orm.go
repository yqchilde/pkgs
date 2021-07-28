package orm

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	gromopentracing "gorm.io/plugin/opentracing"

	"github.com/yqchilde/gint/pkg/logger"
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

// NewMySQL 链接数据库，生成数据库实例
func NewMySQL(c *Config) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		c.Username,
		c.Password,
		c.Addr,
		c.Database,
		true,
		//"Asia/Shanghai"),
		"Local")

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Panicf("open mysql failed. database name: %s, err: %+v", c.Database, err)
	}
	// set for db connection
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(c.ConnMaxLifeTime)

	db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), gormConfig(c))
	if err != nil {
		logger.Panicf("database connection failed. database name: %s, err: %+v", c.Database, err)
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	// set trace
	err = db.Use(gromopentracing.New())
	if err != nil {
		logger.Panicf("using gorm opentracing, err: %+v", err)
	}

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
