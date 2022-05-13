package main

import (
	"fmt"
	"log"
	"wzqb/app"
	"wzqb/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	//
	r.Use(gin.Logger(), gin.Recovery())
	//
	r.SetTrustedProxies([]string{"192.168.20.35"})
	//
	app.RegisterRouter(r) // "The router routes you to a route”
	//
	err := conf.Run()
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		panic(err)
	}
	//
	r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
