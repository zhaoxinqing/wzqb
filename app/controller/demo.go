package controller

import (
	"Kilroy/app/auth"
	"Kilroy/app/common"
	"Kilroy/app/logic"

	"github.com/gin-gonic/gin"
)

// 获取
func TestGet(c *gin.Context) {
	id := c.Query("id")
	common.ResSuccess(c, id)
	// common.ResSuccess(c, nil)
	// common.ResFalse(c, "错误的id参数")
}

// 创建
func TestPost(c *gin.Context) {
	var a auth.A
	err := c.BindJSON(&a)
	if err != nil {
		common.ResFalse(c, common.ErrParam)
		return
	}
	b, err := logic.Demo(a)
	if err != nil {
		common.ResFalse(c, err.Error())
		return
	}
	common.ResSuccess(c, b)
}

// 更新
func TestPut(c *gin.Context) {
	common.ResSuccess(c, "ok")
}

// 删除
func TestDelete(c *gin.Context) {
	common.ResSuccess(c, "ok")
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
