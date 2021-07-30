package middleware

import (
	"Kilroy/common"
	"Kilroy/config"
	"Kilroy/lib"
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		if token := lib.GetHeaderToken(c); token != "" {
			t := strings.ReplaceAll(token, "Bearer ", "")
			isPass, _ := GetTokenFromRedis(t)
			if !isPass {
				common.ResFalse(c, "token无效")
				return
			}
		}
	}
}

// AddTokenToRedis 用户登录，添加账户系统返回的token到redis
func AddTokenToRedis(token string, userIDStr string) (err error) {
	var (
		ctx = context.Background()
	)
	conf := config.GetConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Password,
		DB:       0,
	})
	err = rdb.SetEX(ctx, token, userIDStr, time.Hour*12).Err()
	if err != nil {
		panic(err)
	}
	rdb.Close()
	return
}

// GetTokenFromRedis 从redis中查询时候token是否存在
func GetTokenFromRedis(token string) (isPass bool, err error) {
	var (
		ctx = context.Background()
	)
	conf := config.GetConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Password,
		DB:       0,
	})
	n, err := rdb.Exists(ctx, token).Result()
	if err != nil {
		panic(err)
	}
	isPass = false
	if n > 0 {
		isPass = true
	}
	rdb.Close()
	return
}
