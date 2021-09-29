package app

import (
	"Moonlight/app/controller"

	"github.com/gin-gonic/gin"
)

// Register ...
func Register(app *gin.RouterGroup) {
	var orm = controller.OrmTest{}

	// orm
	app.POST("orm/create", orm.CreateUser)    // 添加新用户
	app.GET("orm/get", orm.GetUsers)          // 获取用户信息
	app.PUT("orm/update", orm.UpdateUser)     // 编辑更新用户信息
	app.DELETE("orm/delete", orm.DeleteUsers) // 抹除用户信息
	app.GET("orm/time", orm.GetByTime)        // 通过时间获取

	// conv
	app.GET("test/struct", controller.StructToStruct)

	//test
	app.GET("test/time", controller.StrTime)

	// upload\download
	app.POST("file/upload", controller.Upload)
	app.GET("file/download", controller.Download)
	app.PUT("file/edit", controller.Edit)

}
