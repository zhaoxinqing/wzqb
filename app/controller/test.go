package controller

import (
	"Kilroy/app/constant"

	"github.com/gin-gonic/gin"
)

// 获取
func TestGet(c *gin.Context) {
	constant.ResSuccess(c, "getAdmin")
}

// 创建
func TestPost(c *gin.Context) {

	param, msg := GetTestPostParam(c)
	if msg != "" {
		constant.ResFail(c, msg)
		return
	}

	a := param.Name + param.Operation
	constant.ResSuccess(c, a)
}

// 更新
func TestPut(c *gin.Context) {
	constant.ResSuccess(c, "putAdmin")
}

// 删除
func TestDelete(c *gin.Context) {
	constant.ResSuccess(c, "deleteAdmin")
}

type PostParam struct {
	Name      string // 名称
	Age       int64  // 年龄
	Operation string // 操作
}

func GetTestPostParam(c *gin.Context) (param PostParam, msg string) {
	err := c.Bind(&param)
	if err != nil {
		msg = "参数解析出错"
		return
	}
	if param.Age == 0 || param.Name == "" || param.Operation == "" {
		msg = "参数传入不完整"
		return
	}
	return
}
