package main

import (
	"Moonlight/app"
	"Moonlight/config"
	"Moonlight/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	//log
	engine := gin.Default()
	engine.Use(middleware.LoggerToFile())
	//
	router := gin.New()
	app.Register(router.Group("moonlight"))
	router.Run(":8990")
}
