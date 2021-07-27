package main

import (
	"Kilroy/app"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	app.InitDB()
	app.RegisterRouting(router.Group("api"))
	router.Run(":8990")
}
