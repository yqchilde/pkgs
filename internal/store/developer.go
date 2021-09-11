package store

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/yqchilde/gin-skeleton/internal/model"
)

type DBDeveloper struct {
	orm *gorm.DB
}

func NewDBDeveloper() *DBDeveloper {
	return &DBDeveloper{orm: DB}
}

func (db *DBDeveloper) Insert(d *model.Developer) error {
	err := db.orm.Create(&d).Error
	if err != nil {
		return errors.Wrap(err, "[store.developer] insert developer err")
	}

	return err
}

func (db *DBDeveloper) GetByEmail(email string) (*model.Developer, error) {
	var developer model.Developer
	err := db.orm.Where("email = ?", email).First(&developer).Error
	if err != nil {
		return nil, errors.Wrap(err, "[repo.user_base] get user err by email")
	}

	return &developer, err
}
