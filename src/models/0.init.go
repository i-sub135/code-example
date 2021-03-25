package models

import "code-example/src/db"

type (
	RestModels struct{}
)

var (
	connect      db.RestDB
	sql          = connect.Sql()
	redisClien   = connect.RedisConnect(0)
	redisClient1 = connect.RedisConnect(1)
)
