package main

import (
	"Kilroy/app"
	"Kilroy/app/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载数据库
	var c model.Conf
	conf := c.GetConf()
	InitDB(conf)
	// 路由注册
	router := gin.New()
	v1 := router.Group("/v1")
	app.RegisterV1(v1)
	app.RegisterV2(v1)
	router.Run(":8990")
}

func InitDB(config *model.Conf) {
	model.InitDB(config) // 初始化数据库
	model.Migration()    // 数据库表迁移（自创建数据库）
}
