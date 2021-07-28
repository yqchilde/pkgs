package model

import (
	"gorm.io/gorm"

	"github.com/yqchilde/gint/pkg/logger"
	"github.com/yqchilde/gint/pkg/storage/orm"
)

var DB *gorm.DB

// Init mysql connection db
func Init(cfg *orm.Config) *gorm.DB {
	DB = orm.NewMySQL(cfg)

	if cfg.AutoMigrate {
		if err := DB.AutoMigrate(
			new(User),
		); err != nil {
			logger.Errorf("gorm auto migrate, err: %+v", err)
		}
	}

	return DB
}

// GetDB return db
func GetDB() *gorm.DB {
	return DB
}
