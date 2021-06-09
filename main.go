package main

import (
	"Kilroy/app/controller"
	"Kilroy/app/models"
	"Kilroy/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var c config.Conf
	conf := c.GetConf()
	fmt.Println(conf)

	InitDB(conf)

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
		v2.POST("/upload", controller.UploadFile)    // 文件上传
		v2.POST("/upload/csv", controller.UploadCSV) // 文件上传
		v2.POST("/upload/sort", controller.SortCSV)  // 分类上传

		// feature分类求和
		v2.POST("/upload/feature", controller.SortFeature) //
		v2.POST("/upload/doc", controller.DocFeature)      //
	}
	router.Run(":8990")
}

func InitDB(config *config.Conf) {
	models.InitDB(config) // 初始化数据库
	models.Migration()    // 数据库表迁移（自创建数据库）
}
