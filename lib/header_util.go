package lib

import "net/http"

// GetHeaderToken 获取header的token
func GetHeaderToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}
