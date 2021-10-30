package middleware

import (
	"Moonlight/app/common"
	"Moonlight/config"
	"context"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// UserClaims ...
type UserClaims struct {
	JobID    string // 工号
	UserName string // 用户名
	jwt.StandardClaims
}

// GenerateToken 。。。
func GenerateToken(jobID string, userName string) (string, error) {
	expires := 12
	claims := UserClaims{
		JobID:    jobID,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expires)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("golang")) // 签名密钥
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken。。。
func ParseToken(bearerToken string) (*UserClaims, error) {
	token := strings.ReplaceAll(bearerToken, "Bearer ", "")
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

const (
	token_key = "Authorization" //页面token键名
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		if token := c.GetHeader(token_key); token != "" {
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
		ctx  = context.Background()
		conf = config.GetConfig()
	)
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
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
		ctx  = context.Background()
		conf = config.GetConfig()
	)
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
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
