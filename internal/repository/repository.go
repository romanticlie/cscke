package repository

import "cscke/pkg/db"

var (
	D, _  = db.MysqlConnect()
	Rd, _ = db.RedisConnect()
)
