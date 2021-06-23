package app

import (
	"Kilroy/app/controller"

	"github.com/gin-gonic/gin"
)

// RegisterRouterSys ...
func RegisterV1(v1 *gin.RouterGroup) {
	// 个人中心
	v1.GET("/someGet", controller.TestGet)          // 获取
	v1.POST("/somePost", controller.TestPost)       // 创建
	v1.PUT("/somePut", controller.TestPut)          // 更新
	v1.DELETE("/someDelete", controller.TestDelete) // 删除
}

// RegisterRouterSys ...
func RegisterV2(v2 *gin.RouterGroup) {
	v2.POST("/upload", controller.UploadFile)          // 文件上传
	v2.POST("/upload/csv", controller.UploadCSV)       // 文件上传
	v2.POST("/upload/sort", controller.SortCSV)        // 分类上传
	v2.POST("/upload/feature", controller.SortFeature) //
	v2.POST("/upload/doc", controller.DocFeature)      // 分类统计

}
