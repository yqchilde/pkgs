package dao

import (
	"gorm.io/gorm"

	"github.com/yqchilde/gint/pkg/config"
)

var AuthDao *Dao

type Dao struct {
	db  *gorm.DB
	cfg *config.Config
	// todo cache
	// todo tracer
}

func New(db *gorm.DB, cfg *config.Config) *Dao {
	d := &Dao{
		db:  db,
		cfg: cfg,
	}

	AuthDao = d
	return d
}
