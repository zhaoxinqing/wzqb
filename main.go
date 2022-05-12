package main

import (
	"go-template/wzqb"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	// conf := GetConfigInformation("config.yanml")
	// fmt.Println(conf)
	// Creates a router without any middleware by default
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.20.35"})
	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", wzqb.GetAllMenu)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/api")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use()
	{
		authorized.POST("/login", wzqb.Login)                  // login
		authorized.GET("/array_to_string", wzqb.ArrayToString) // ArrayToString

		// nested group
		testing := authorized.Group("/data")
		testing.GET("/analytics", wzqb.ArrayToString)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
