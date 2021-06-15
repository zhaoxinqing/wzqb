package constant

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseModel ...
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseModelBase ...
type ResponseModelBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{Code: SUCCESS_CODE, Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// ResSuccessMsg 响应成功
func ResSuccessMsg(c *gin.Context) {
	ret := ResponseModelBase{Code: SUCCESS_CODE, Message: "ok"}
	ResJSON(c, http.StatusOK, &ret)
}

// ResFail 响应失败
func ResFail(c *gin.Context, msg string) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// ResFailCode 响应失败
func ResFailCode(c *gin.Context, msg string, code int) {
	ret := ResponseModelBase{Code: code, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
	c.Abort()
}

// ResErrSrv 响应错误-服务端故障
func ResErrSrv(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: err.Error()}
	ResJSON(c, http.StatusOK, &ret)
}

// ResErrCli 响应错误-用户端故障
func ResErrCli(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: FAIL_CODE, Message: err.Error()}
	ResJSON(c, http.StatusOK, &ret)
}

// ResponsePageData ...
type ResponsePageData struct {
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}

// ResponsePage ...
type ResponsePage struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponsePageData `json:"data"`
}

// ResSuccessPage 响应成功-分页数据
func ResSuccessPage(c *gin.Context, total int64, list interface{}) {
	ret := ResponsePage{Code: SUCCESS_CODE, Message: "ok", Data: ResponsePageData{Total: total, Items: list}}
	ResJSON(c, http.StatusOK, &ret)
}
