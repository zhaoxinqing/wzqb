package main

import (
	"Kilroy/app"
	"Kilroy/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.InitDB()
	app.Register(router.Group("api"))
	router.Run(":8990")
}
