package controller

import (
	"Kilroy/app/common"
	"Kilroy/app/model"

	"github.com/gin-gonic/gin"
)

type TestDB struct{}

// 添加
func (i TestDB) Add(c *gin.Context) {
	info := model.User{}
	// 入参校验
	err := c.BindJSON(&info)
	if err != nil {
		common.ResFalse(c, common.ErrParam)
		return
	}
	if err := model.Create(&info); err != nil {
		common.ResFalse(c, err.Error())
	}
	common.ResSuccess(c, info)
}

// 更新
func (i TestDB) Update(c *gin.Context) {
	info := model.User{}
	err := c.BindJSON(&info)
	if err != nil {
		common.ResFalse(c, common.ErrParam)
		return
	}
	err = model.DB.Model(&model.User{}).Where("id=?", info.ID).Update("remark", info.Remark).Error
	if err != nil {
		common.ResFalse(c, common.ErrDB)
		return
	}
	common.ResSuccess(c, nil)
}
