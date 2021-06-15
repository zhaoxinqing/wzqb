package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func Export() {
	//表头
	titleList := []string{"ID", "用户", "IP地址", "登陆时间", "说明"}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		var data []interface{}
		//调用方法
		ExportToExcel(c, titleList, data, "xx导出")
	})
	_ = r.Run()
}

//数据导出excel并下载
func ExportToExcel(c *gin.Context, titleList []string, data []interface{}, fileName string) {
	// 生成一个新的文件
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")
	// 插入表头
	titleRow := sheet.AddRow()
	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
		//表头字体颜色
		cell.GetStyle().Font.Color = "00FF0000"
		//居中显示
		cell.GetStyle().Alignment.Horizontal = "center"
		cell.GetStyle().Alignment.Vertical = "center"
	}
	// 插入内容
	for _, v := range data {
		row := sheet.AddRow()
		row.WriteStruct(v, -1)
	}
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename=\"%s-%s.xlsx\"", fileName, time.Now().Format("2006-01-02 15:04:05"))
	c.Writer.Header().Set("Content-Disposition", disposition)
	_ = file.Write(c.Writer)
}
