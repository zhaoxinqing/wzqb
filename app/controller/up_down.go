package controller

import (
	"Moonlight/app/common"
	"Moonlight/app/logic"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//
func Upload(c *gin.Context) {
	var (
		filePrefix = fmt.Sprintf("%d", time.Now().Unix())
		file, _    = c.FormFile("file")
		filePath   = "./docs/upload/" + filePrefix + file.Filename
	)
	_ = c.SaveUploadedFile(file, filePath)
	common.ResSuccess(c, nil)
}

//
func Download(c *gin.Context) {
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "163.pdf")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./docs/upload/163.pdf")

}

//
func Edit(c *gin.Context) {
	var (
		path    = "./docs/102.csv"
		newPath = "./docs/102的副本.csv"
	)
	logic.OperateCSV(path, newPath)
	common.ResSuccess(c, nil)
}
