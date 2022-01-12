package model

const (
	GenderNone = 0 //性别无
	GenderMale = 1 //性别男
	GenderFemale = 2//性别女

	StatusNormal = 0 //状态：正常
	StatusForbidden = 1//状态：封禁
)

type User struct {
	Userid uint64 `gorm:"type:bigint unsigned;primaryKey;comment:用户id"`
	InviteId uint64 `gorm:"type:bigint unsigned;not null;default:0;index:invite_id;comment:邀请者用户id"`
	Nickname string `gorm:"size:32;not null;index:nickname;comment:用户昵称"`
	Gender int `gorm:"type:tinyint;not null;default:0;comment:性别"`
	Avatar string `gorm:"size:255;null;comment:头像"`
	Telephone string `gorm:"size:11;null;unique;comment:手机号"`
	Password string `gorm:"size:255;null;comment:密码"`
	Status int `gorm:"type:tinyint;not null;default:0;comment:状态"`
	Email string `gorm:"size:64;null;comment:邮箱"`
	BirthYear int `gorm:"type:int;null;comment:出生年份"`
	BirthMonth int `gorm:"type:int;null;comment:出生月份"`
	BirthDay int `gorm:"type:int;null;comment:出生日期"`
	Province int `gorm:"type:int;null;comment:省份"`
	City int `gorm:"type:int;null;comment:城市"`
	County int `gorm:"type:int;null;comment:区县"`
	RegisterIp string `gorm:"size:64;not null;default:'';comment:注册IP"`
	MysqlTimestamps
}


func (u User) TableName() string{
	return "user"
}