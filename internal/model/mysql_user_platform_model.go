package model

const (
	PlatFormWechatWeb = 1 + iota
	PlatFormWechatGZH
	PlatFormWechatApp
)

type UserPlatform struct {
	Userid   uint64 `gorm:"type:bigint unsigned;primaryKey;comment:用户id"`
	Platform int    `gorm:"type:tinyint;not null;default:0;uniqueIndex:platform_openid;comment:平台类型"`
	Openid   string `gorm:"size:64;not null;uniqueIndex:platform_openid;comment:平台openid"`
	UnionId  string `gorm:"size:100;null;comment:平台唯一的union_id"`
	MysqlTimestamps
}

func (u UserPlatform) TableName() string {
	return "user_platform"
}
