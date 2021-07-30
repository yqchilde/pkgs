package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/yqchilde/gint/pkg/validator"
)

type Tabler interface {
	TableName() string
}

type BaseModel struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

var (
	v            = validator.NewValidator()
	CasbinVerify = validator.Rules{
		"AuthorityID": {v.NotEmpty()},
		"CasbinInfos": {v.NotEmpty(), v.Gt("0")},
	}
	SignUpVerify = validator.Rules{
		"Email":       {v.NotEmpty()},
		"Username":    {v.NotEmpty()},
		"Password":    {v.NotEmpty()},
		"AuthorityID": {v.NotEmpty()},
	}
	SignInVerify = validator.Rules{
		"Email":    {v.NotEmpty()},
		"Password": {v.NotEmpty()},
	}
)
