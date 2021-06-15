package main

import (
	"Kilroy/app"
	"Kilroy/app/models"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// 加载数据库
	var c models.Conf
	conf := c.GetConf()
	fmt.Println(conf)
	InitDB(conf)

	// 路由注册
	router := gin.New()
	v1 := router.Group("/v1")
	app.RegisterV1(v1)
	app.RegisterV2(v1)
	router.Run(":8990")
}

func InitDB(config *models.Conf) {
	models.InitDB(config) // 初始化数据库
	models.Migration()    // 数据库表迁移（自创建数据库）
}
