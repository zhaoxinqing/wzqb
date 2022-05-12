package main

import (
	"go-template/app"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	// conf := GetConfigInformation("config.yanml")
	// fmt.Println(conf)
	// Creates a router without any middleware by default
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.20.35"})

	app.RegisterRouter(r) // "The router routes you to a route”

	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
