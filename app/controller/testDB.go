package ctrl

import (
	"Kilroy/app/common"
	"Kilroy/app/model"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

type TestDB struct{}

type TestAddParam struct {
	Name   string `json:"name"`   // 用户名
	Phone  string `json:"phone"`  // 用户手机号
	Status int64  `json:"status"` // 用户状态（1:正常 2:未激活 3:暂停使用）
	Beat   string `json:"beat"`
}

//
func (i TestDB) Add(c *gin.Context) {
	var param TestAddParam
	// 入参校验
	err := c.BindJSON(&param)
	if err != nil {
		common.ResFalse(c, common.ErrParam)
		return
	}
	info := new(model.Users)
	_ = gconv.Struct(param, &info)
	if err := model.Create(&info); err != nil {
		common.ResFalse(c, err.Error())
	}
	common.ResSuccess(c, info)
}
