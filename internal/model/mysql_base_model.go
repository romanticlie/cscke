package model

import (
	"time"
)

// MysqlPrimaryKey default 主键
type MysqlPrimaryKey struct {
	Id uint64 `gorm:"type:bigint unsigned;not null;primaryKey;autoIncrement:true;comment:ID主键"`
}

// MysqlTimestamps 创建和更新的字段
type MysqlTimestamps struct {
	CreatedAt time.Time `gorm:"type:timestamp;index:created_at;comment:创建日期"`
	UpdatedAt time.Time `gorm:"type:timestamp;comment:修改日期"`
}
