package main

import (
	"Kilroy/app"
	"Kilroy/app/model"

	"github.com/gin-gonic/gin"
)

func main() {
	var c model.Conf
	conf := c.GetConf()
	InitDB(conf)
	router := gin.New()
	app.RegisterRouting(router.Group("/v1"))
	router.Run(":8990")
}

func InitDB(config *model.Conf) {
	model.InitDB(config) // 初始化数据库
	model.Migration()    // 数据库表迁移（自创建数据库）
}
