package app

import (
	"go-template/wzqb"

	"github.com/gin-gonic/gin"
)

// RegisterRouter ...
func RegisterRouter(r *gin.Engine) {
	r.GET("/benchmark", wzqb.GetAllMenu)
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/api")
	// Per route middleware, you can add as many as you desire.
	r.Use()
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
}
