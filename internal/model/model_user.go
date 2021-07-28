package model

type User struct {
	UserID   string `json:"user_id" gorm:"primaryKey;type:char(36)"`
	Username string `json:"username"`
	Password string `json:"-"`

	BaseModel
}

func (User) TableName() string {
	return "user"
}
