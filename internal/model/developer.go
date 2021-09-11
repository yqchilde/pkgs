package model

type Developer struct {
	UserID    string `json:"user_id,omitempty" gorm:"primaryKey;type:char(36)"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      int    `json:"role"`

	BaseModel
}

func (u *Developer) TableName() string {
	return "developer"
}
