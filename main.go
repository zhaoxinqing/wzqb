package main

import (
	"Kilroy/app/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个不包含中间件的路由器
	router := gin.New()

	// Simple group: v1
	app := router.Group("/v1")
	{
		app.GET("/someGet", controller.TestGet)          // 获取
		app.POST("/somePost", controller.TestPost)       // 创建
		app.PUT("/somePut", controller.TestPut)          // 更新
		app.DELETE("/someDelete", controller.TestDelete) // 删除
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/upload", controller.UploadFile) // 文件上传
	}

	router.Run(":8900")
}
