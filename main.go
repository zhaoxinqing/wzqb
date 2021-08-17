package main

import (
	"Kilroy/app"
	"Kilroy/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.InitDB()
	// service.TimeToNotification() // 定时任务
	app.Register(router.Group("api"))
	router.Run(":8990")
}
