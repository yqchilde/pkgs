package model

type Application struct {
	AppID     string `json:"app_id,omitempty" gorm:"primaryKey;type:char(36)"`
	AppName   string `json:"app_name"`
	Creator   string `json:"creator" gorm:"type:char(36);index"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Status    uint32 `json:"status" gorm:"type:tinyint(1);default:0"`

	BaseModel
}

func (u *Application) TableName() string {
	return "application"
}
