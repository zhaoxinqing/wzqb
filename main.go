package main

import (
	"Kilroy/app/controller"
	"Kilroy/config"
	"Kilroy/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	msg := "测试打印耗时："
	utils.GetTime(msg)
	// 加载配置
	configPath := "config/env.yaml"
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	initDB(conf)

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
		// v2.POST("/upload/csv", controller.UploadCSV) // 文件上传
		v2.POST("/upload/csv", controller.SortCSV) // 分类上传

	}

	router.Run(":8990")
}

func initDB(config *config.Config) {

	// models.InitDB(config)
	// models.Migration()
}
