package model

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	ID        uint         `json:"-" gorm:"primarykey"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}

type Tabler interface {
	TableName() string
}
