package app

import (
	ctrl "Kilroy/app/controller"

	"github.com/gin-gonic/gin"
)

// RegisterRouterSys ...
func RegisterRouting(app *gin.RouterGroup) {
	// 个人中心
	demo := ctrl.Demo{}
	app.GET("/someGet", demo.TestGet)                    // 获取
	app.POST("/somePost", demo.TestPost)                 // 创建
	app.PUT("/somePut", demo.TestPut)                    // 更新
	app.DELETE("/someDelete", demo.TestPut)              // 删除
	app.POST("/test/upload/table", demo.UploadTable)     //
	app.POST("/test/upload/html_to_pdf", demo.HtmlToPDF) //

	db := ctrl.TestDB{}
	app.POST("/test_db/add", db.Add)
	// 文件
	file := ctrl.File{}
	app.POST("/upload", file.UploadFile)          // 文件上传
	app.POST("/upload/csv", file.UploadCSV)       // 文件上传
	app.POST("/upload/sort", file.SortCSV)        // 分类上传
	app.POST("/upload/feature", file.SortFeature) //
	app.POST("/upload/doc", file.FindSH)          //
}
