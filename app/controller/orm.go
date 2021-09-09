package controller

import (
	"Kilroy/app/common"
	"Kilroy/app/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// CreateUser 添加
func (i UserController) CreateUser(c *gin.Context) {
	var (
		info = model.User{}
		aa   = model.User{}
		err  error
	)
	// 获取入参及赋值
	if err = c.BindJSON(&info); err != nil {
		common.ResFalse(c, common.ErrParam)
		return
	}
	//
	err = model.DB.Where("id = 6").Find(&aa).Error
	aa.ID = 0
	if err = model.Create(&aa); err != nil {
		common.ResFalse(c, err.Error())
		return
	}
	common.ResSuccess(c, info)
}

// GetUsers 获取
func (i UserController) GetUsers(c *gin.Context) {
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
func (i UserController) UpdateUser(c *gin.Context) {
	var (
		info = model.User{}
		err  error
	)
	// 入参校验
	if err = c.BindJSON(&info); err != nil {
		common.ResFalse(c, common.ErrParam)
		return
	}
	if err = model.Save(&info); err != nil {
		common.ResFalse(c, err.Error())
		return
	}
	common.ResSuccess(c, info)
}

// DeleteUsers 删除
func (i UserController) DeleteUsers(c *gin.Context) {
	var (
		err error
	)
	//  获取id参数
	idStr, get := c.GetQuery("id")
	if !get {
		common.ResFalse(c, "获取id参数失败")
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResFalse(c, "参入参数id格式不正确")
		return
	}

	if err := model.DB.Delete(&model.User{ID: id}); err != nil {
		common.ResFalse(c, "")
		return
	}
	common.ResSuccess(c, nil)
}

type GetByTimeParam struct {
	Time1 string `json:"time1"`
	Time2 string `json:"time2"`
}

// GetByTime 通过时间获取
func (i UserController) GetByTime(c *gin.Context) {
	var (
		db    = model.DB
		infos = []model.User{}
		err   error
	)
	// 入参校验
	time1, _ := c.GetQuery("time1") // 2021-08-20
	time2, _ := c.GetQuery("time2") // 2021-08-28

	// time.Sleep(1 * time.Second)

	err = db.Debug().Model(&model.User{}).Where("updated_at between ? and ? ", time1, time2).
		Find(&infos).Error
	if err != nil {
		common.ResFalse(c, err.Error())
	}
	common.ResSuccess(c, infos)
}
