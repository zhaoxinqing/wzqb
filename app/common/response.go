package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResSucModel struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResFalModel struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, data interface{}) {
	ret := ResSucModel{Code: SUCCESS_CODE, Message: "ok", Data: data}
	ResponseJSON(c, http.StatusOK, &ret)
}

// ResFalse 响应失败
func ResFalse(c *gin.Context, msg string) {
	ret := ResFalModel{Code: FALSE_CODE, Message: msg}
	ResponseJSON(c, http.StatusOK, &ret)
}

// 响应JSON数据
func ResponseJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
	c.Abort()
}
