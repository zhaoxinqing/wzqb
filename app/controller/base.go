package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// GetPageSort 获取排序信息
func GetPageSort(c *gin.Context) string {
	return GetQueryToStr(c, "sort")
}

// GetPageKey 获取搜索关键词信息
func GetPageKey(c *gin.Context) string {
	return GetQueryToStr(c, "key")
}

// GetQueryToStrE ...
func GetQueryToStrE(c *gin.Context, key string) (string, error) {
	str, ok := c.GetQuery(key)
	if !ok {
		return "", errors.New("没有这个值传入")
	}
	return str, nil
}

// GetQueryToStr ...
func GetQueryToStr(c *gin.Context, key string, defaultValues ...string) string {
	var defaultValue string
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	str, err := GetQueryToStrE(c, key)
	if str == "" || err != nil {
		return defaultValue
	}
	return str
}

// // GetPageIndex 获取页码
// func GetPageIndex(c *gin.Context) int64 {
// 	return GetQueryToInt64(c, "page", 1)
// }

// // GetPageSize 获取每页记录数
// func GetPageSize(c *gin.Context) int64 {
// 	limit := GetQueryToInt64(c, "pageSize", 10)
// 	if limit > 500 {
// 		limit = 20
// 	}
// 	return limit
// }
