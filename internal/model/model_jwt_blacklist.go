package model

type JwtBlacklist struct {
	ID  uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Jwt string `json:"jwt" gorm:"type:text"`

	BaseModel
}

func (JwtBlacklist) TableName() string {
	return "jwt_blacklist"
}
