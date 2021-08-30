package main

import (
	"Kilroy/app"
	"Kilroy/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	config.InitDB()

	// 注册路由
	router := gin.New()
	app.Register(router.Group("moonlight"))
	router.Run(":8990")
}
