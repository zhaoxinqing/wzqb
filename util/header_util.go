package util

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// 获取 Header 的Token
func GetHeaderToken(c *gin.Context) string {
	return c.Request.Header.Get("Authorization")
}

// 获取 Query 查询参数
func GetQuery(c *gin.Context, key string) (string, error) {
	str, ok := c.GetQuery(key)
	if !ok {
		return "", errors.New("获取" + key + "参数失败")
	}
	return str, nil
}

// 获取 Header 信息
func GetHeader(c *gin.Context, key string) string {
	return c.Request.Header.Get(key)
}
