package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	SlowThreshold   time.Duration `mapstructure:"slow-threshold"`
}

// NewMySQL generate mysql orm instance
func NewMySQL(c *Config) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		c.Username,
		c.Password,
		c.Addr,
		c.Database,
		true,
		"Local",
	)

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

	log.Println("mysql connected and running!")
	return db
}

func gormConfig(c *Config) *gorm.Config {
	// Foreign key constraints are prohibited
	// Foreign key constraints are not recommended for production environments
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}

	// Print all logs
	if c.ShowLog {
		config.Logger = logger.Default.LogMode(logger.Info)
	} else {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}

	// Print slow query log
	if c.SlowThreshold > 0 {
		config.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: c.SlowThreshold,
				Colorful:      true,
				LogLevel:      logger.Warn,
			},
		)
	}
	return config
}
