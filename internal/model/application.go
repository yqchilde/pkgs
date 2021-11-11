package model

type Application struct {
	BaseModel
	AppID     string `json:"app_id,omitempty" gorm:"type:char(36);index"`
	AppName   string `json:"app_name"`
	Creator   string `json:"creator" gorm:"type:char(36);index"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Status    uint32 `json:"status" gorm:"type:tinyint(1);default:0"`
}

func (u *Application) TableName() string {
	return "application"
}
