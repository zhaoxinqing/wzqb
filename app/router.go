package app

import (
	"wzqb/wzqb"

	"github.com/gin-gonic/gin"
)

// RegisterRouter ...
func RegisterRouter(r *gin.Engine) {
	r.GET("/benchmark", wzqb.GetAllMenu)
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	api := r.Group("/api")
	// Per route middleware, you can add as many as you desire.
	r.Use()
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	api.Use()
	{
		user := api.Group("/user")
		user.POST("/login", wzqb.Login)                  // login
		user.GET("/array_to_string", wzqb.ArrayToString) // ArrayToString

		// nested group
		data := api.Group("/data")
		data.GET("/analytics", wzqb.ArrayToString)
	}
}
