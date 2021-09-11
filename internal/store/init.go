package store

import (
	"log"

	"gorm.io/gorm"

	"github.com/yqchilde/gin-skeleton/internal/model"
	"github.com/yqchilde/gin-skeleton/pkg/storage/orm"
)

var DB *gorm.DB

func Init(cfg *orm.Config) *gorm.DB {
	DB = orm.NewMySQL(cfg)

	if cfg.AutoMigrate {
		if err := DB.AutoMigrate(
			new(model.Developer),
			new(model.Application),
		); err != nil {
			log.Printf("gorm auto migrate, err: %+v", err)
		}
	}

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
