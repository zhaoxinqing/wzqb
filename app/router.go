package app

import (
	"Kilroy/app/controller"

	"github.com/gin-gonic/gin"
)

// RegisterRouterSys ...
func Register(app *gin.RouterGroup) {

	demo := controller.Demo{}
	file := controller.File{}
	db := controller.TestDB{}

	// 个人中心
	app.GET("/someGet", demo.TestGet)                    // 获取
	app.POST("/somePost", demo.TestPost)                 // 创建
	app.PUT("/somePut", demo.TestPut)                    // 更新
	app.DELETE("/someDelete", demo.TestPut)              // 删除
	app.POST("/test/upload/table", demo.UploadTable)     //
	app.POST("/test/upload/html_to_pdf", demo.HtmlToPDF) //

	app.POST("/test_db/add", db.Add)         //
	app.PUT("/test_db/update", db.Update)    //
	app.DELETE("/test_db/delete", db.Update) //

	app.POST("/upload", file.UploadFile)          // 文件上传
	app.POST("/upload/csv", file.UploadCSV)       // 文件上传
	app.POST("/upload/sort", file.SortCSV)        // 分类上传
	app.POST("/upload/feature", file.SortFeature) //
	app.POST("/upload/doc", file.FindSH)          //

	app.GET("/test", controller.TestJson)
}
