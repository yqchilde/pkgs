package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/yqchilde/gint/internal/model"
)

func (d *Dao) GetByEmail(ctx context.Context, email string) (bool, *model.User, error) {
	var user model.User
	err := d.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil, nil
		}
		return false, nil, errors.Wrap(err, "[Repo.public] GetByEmail failed")
	}

	return true, &user, err
}

func (d *Dao) CreateUser(ctx context.Context, u *model.User) error {
	err := d.db.Create(&u).Error
	if err != nil {
		return errors.Wrap(err, "[Repo.public] CreateUser failed")
	}

	return nil
}
