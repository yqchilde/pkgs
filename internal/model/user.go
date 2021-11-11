package model

type User struct {
	BaseModel
	UserID    string `json:"user_id,omitempty" gorm:"index;type:char(36)"`
	AppID     string `json:"app_id" gorm:"index;type:char(36)"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Head      string `json:"head"`
	Online    bool   `json:"online,omitempty" gorm:"default:0"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetUserName() string {
	return u.FirstName + " " + u.LastName
}
