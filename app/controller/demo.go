package controller

import (
	"Kilroy/app/auth"
	"Kilroy/app/common"
	"Kilroy/app/logic"
	"Kilroy/util"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取
func TestGet(c *gin.Context) {
	id := c.Query("id")
	common.ResSuccess(c, id)
	// common.ResSuccess(c, nil)
	// common.ResFalse(c, "错误的id参数")
}

// 创建
func TestPost(c *gin.Context) {
	var a auth.A
	// 入参解析校验
	err := c.BindJSON(&a)
	if err != nil {
		common.ResFalse(c, common.ErrParam)
		return
	}
	// 逻辑处理
	b, err := logic.Demo(a)
	if err != nil {
		common.ResFalse(c, err.Error())
		return
	}
	common.ResSuccess(c, b)
}

// 更新
func TestPut(c *gin.Context) {
	common.ResSuccess(c, "ok")
}

// 删除
func UploadTable(c *gin.Context) {
	file, err := c.FormFile("table")
	if err != nil {
		common.ResFalse(c, err.Error()+"（获取上传文件失败）")
		return
	}
	// 判断文件格式
	fileType := util.ExtractFileType(file.Filename)
	if fileType != "xlsx" {
		common.ResFalse(c, "请上传xlsx表格格式")
		return
	}
	// 保存临时文件
	err = os.MkdirAll("doc/temp", 0777)
	filePath := "doc/temp/" + fmt.Sprintf("%d_", time.Now().Unix()) + file.Filename
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		common.ResFalse(c, err.Error())
		return
	}
	// 执行逻辑
	result, err := util.ReadExcel(filePath)
	common.ResSuccess(c, result)
}

// 删除
func HtmlToPDF(c *gin.Context) {
	file, err := c.FormFile("html")
	if err != nil {
		common.ResFalse(c, err.Error()+"（获取上传文件失败）")
		return
	}
	// 判断文件格式
	fileType := util.ExtractFileType(file.Filename)
	if fileType != "html" {
		common.ResFalse(c, "请上传正确的文件格式")
		return
	}
	// 保存临时文件
	err = os.MkdirAll("doc/temp", 0777)
	filePath := "doc/temp/" + fmt.Sprintf("%d_", time.Now().Unix()) + file.Filename
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		common.ResFalse(c, err.Error())
		return
	}
	// 执行逻辑
	result, err := util.HTMLtoPDF(filePath)
	common.ResSuccess(c, result)
}

type PostParam struct {
	Name      string // 名称
	Age       int64  // 年龄
	Operation string // 操作
}

func GetTestPostParam(c *gin.Context) (param PostParam, msg string) {
	err := c.Bind(&param)
	if err != nil {
		msg = "参数解析出错"
		return
	}
	if param.Age == 0 || param.Name == "" || param.Operation == "" {
		msg = "参数传入不完整"
		return
	}
	return
}
