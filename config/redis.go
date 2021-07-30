package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client
var ctx = context.Background()
var instance *Redis

func InitRedis() {
	conf := GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Database.Host,
		Password: conf.Database.Port,
		DB:       0,
		PoolSize: 132,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("Redis 连接失败 【错误】：%s", err.Error())
		return
	}
	fmt.Printf("Redis 连接成功： %s", pong)
}
