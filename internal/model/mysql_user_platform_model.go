package model


const (
	TypeWechatApp = 1 + iota
	TypeWechatGZH
	TypeWechatWeb
	TypeQQ
	TypeWeibo
	TypeApple
)

type UserPlatform struct {
	Userid uint64 `gorm:"type:bigint unsigned;primaryKey;comment:用户id"`
	Type int `gorm:"type:tinyint;not null;default:0;comment:平台类型"`
	Openid string `gorm:"size:64;not null;comment:平台openid"`
	UnionId string `gorm:"size:100;null;comment:平台唯一的union_id"`
	MysqlTimestamps
}

func (u UserPlatform) TableName() string {
	return "user_platform"
}