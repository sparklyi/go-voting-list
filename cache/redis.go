package cache

import (
	"context"
	"gin_Ranking/config"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	//初始化连接
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPasswd,
		DB:       config.RedisDB,
	})
	Rctx = context.Background()

}
