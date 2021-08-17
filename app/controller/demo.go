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

type Demo struct{}

// 获取
func (i Demo) TestGet(c *gin.Context) {
	id := c.Query("id")
	common.ResSuccess(c, id)
	// common.ResSuccess(c, nil)
	// common.ResFalse(c, "错误的id参数")
}

// 创建
func (i Demo) TestPost(c *gin.Context) {
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

type UploadMessage struct {
	Message string `json:"message"`
}

// 更新
func (i Demo) TestPut(c *gin.Context) {
	var (
		records     []string
		successList = []interface{}{}
		errorMap    = make(map[string]string)
	)
	message := UploadMessage{
		Message: fmt.Sprintf("共 %d 条数据，成功添加 %d 条，错误 %d 条", len(records)-2, len(successList), len(errorMap)),
	}
	common.ResSuccess(c, message)
}

// 删除
func (i Demo) UploadTable(c *gin.Context) {
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
func (i Demo) HtmlToPDF(c *gin.Context) {
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

func (i Demo) GetTestPostParam(c *gin.Context) (param PostParam, msg string) {
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
