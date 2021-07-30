package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/yqchilde/gint/internal/model"
)

func (d *Dao) GetByJwtToken(ctx context.Context, jwtToken string) (bool, *model.JwtBlacklist, error) {
	var jwt model.JwtBlacklist
	err := d.db.Where("jwt = ?", jwtToken).First(&jwt).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil, nil
		}
		return false, nil, errors.Wrap(err, "[Repo.jwt_blacklist] GetByJwtToken failed")
	}

	return true, &jwt, nil
}

func (d *Dao) AddJwtBlacklist(ctx context.Context, jwtBlacklist *model.JwtBlacklist) error {
	err := d.db.Create(&jwtBlacklist).Error
	if err != nil {
		return errors.Wrap(err, "[Repo.jwt_blacklist] AddJwtBlacklist failed")
	}

	return nil
}
