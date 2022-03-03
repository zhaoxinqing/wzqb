package common

import (
	"net/http"
)

// 常用定义
const (
	SUCCESS_CODE         = 200                          // 成功的状态码
	FALSE_CODE           = 10000                        // 失败的状态码
	EmptyID              = 0                            // 空id
	ErrParam             = "err param format"           // 请求参数格式不正确
	ParamMissingErr      = "missing required parameter" // 缺少必要参数
	AuthenticationFailed = "user authentication failed" // 用户token验证失败
)

// 、
const (
	IncorrectParameteFormat  = "参数格式不正确,请重试"
	MissingRequiredParameter = "参数缺失,请重试"
)

// GetHeaderToken 获取header的token
func GetHeaderToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}

// IsContainInt64 ...
func IsContainNumber(items []int64, item int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// IsContainString ...
func IsContainString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// RemoveDupNumber ...
func RemoveDupNumber(items []int64) []int64 {
	result := make([]int64, 0, len(items))
	tempMap := map[int64]byte{}
	for _, eachItem := range items {
		l := len(tempMap)
		tempMap[eachItem] = 0
		if len(tempMap) != l {
			result = append(result, eachItem)
		}
	}
	return result
}
