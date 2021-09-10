package main

import (
	"Kilroy/app"
	"Kilroy/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	router := gin.New()
	app.Register(router.Group("moonlight"))
	router.Run(":8990")
}
