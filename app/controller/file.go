package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("upload")
	filename := file.Filename
	filePath := "./docs/upload/" + fmt.Sprintf("%d", time.Now().Unix()) + filename
	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
	}
	c.String(200, "Success")
	// common.ResSuccess(c, "postAdmin")
}
