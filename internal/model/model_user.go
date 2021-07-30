package model

type User struct {
	UserID      string `json:"user_id" gorm:"primaryKey;type:char(36)"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	AuthorityID string `json:"authority_id"`

	BaseModel
}

func (User) TableName() string {
	return "user"
}
