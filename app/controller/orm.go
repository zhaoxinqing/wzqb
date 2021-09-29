package controller

import (
	"Moonlight/app/common"
	"Moonlight/app/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrmTest struct{}

// CreateUser 添加
func (*OrmTest) CreateUser(c *gin.Context) {
	var (
		info = model.User{}
		err  error
	)
	// 获取入参及赋值
	if err = c.BindJSON(&info); err != nil {
		common.ResFalse(c, common.ErrParam)
		return
	}
	//
	err = model.DB.Where("id = 6").UpdateColumn("id", 7).Error
	if err != nil {
		common.ResFalse(c, err.Error())
		return
	}
	common.ResSuccess(c, info)
}

// GetUsers 获取
func (*OrmTest) GetUsers(c *gin.Context) {
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
func (*OrmTest) UpdateUser(c *gin.Context) {
	var (
		info = model.User{}
		err  error
	)
	if err = c.BindJSON(&info); err != nil {
		common.ResFalse(c, common.ErrParam)
		return
	}
	err = model.DB.Model(&model.User{}).Where("id=?", info.ID).Update("remark", info.Remark).Error
	if err != nil {
		common.ResFalse(c, err.Error())
		return
	}
	common.ResSuccess(c, info)
}

// DeleteUsers 删除
func (*OrmTest) DeleteUsers(c *gin.Context) {
	var err error

	//  获取id参数
	idStr, _ := c.GetQuery("id")

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

// GetByTime 通过时间获取
func (*OrmTest) GetByTime(c *gin.Context) {
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
