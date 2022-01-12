package repository

import "cscke/pkg/db"

var (
	d,_ = db.MysqlConnect()
	rd,_ = db.RedisConnect()
)
