package store

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/yqchilde/gin-skeleton/internal/model"
)

type DBApplication struct {
	orm *gorm.DB
}

func NewDBApplication() *DBApplication {
	return &DBApplication{orm: DB}
}

func (db *DBApplication) Insert(d *model.Application) error {
	err := db.orm.Create(&d).Error
	if err != nil {
		return errors.Wrap(err, "[store.application] insert application err")
	}

	return err
}

func (db *DBApplication) DeleteByAppID(AppID string) error {
	err := db.orm.Where("app_id = ?", AppID).Delete(&model.Application{}).Error
	if err != nil {
		return errors.Wrap(err, "[store.application] delete by app id err")
	}

	return err
}

func (db *DBApplication) VerifyAppKeyAppSecret(appKey, appSecret string) (bool, error) {
	var app model.Application
	result := db.orm.
		Where("app_key = ? AND app_secret = ?", appKey, appSecret).
		Limit(1).Find(&app)

	if result.Error != nil {
		return false, errors.Wrap(result.Error, "[store.application] verify appKey and appSecret err")
	}
	if result.RowsAffected > 0 {
		return true, nil
	}

	return false, nil
}
