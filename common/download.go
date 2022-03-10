package common

import "github.com/tealeg/xlsx"

// 下载 excel 文件

// DownloadEXcelFiles
func DownloadEXcelFiles(records [][]string, title []string, path string) error {
	// 生成一个新的文件
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")
	// 插入表头
	titleRow := sheet.AddRow()
	for _, v := range title {
		titleRow.AddCell().Value = v
	}
	// 插入内容
	for _, record := range records {
		row := sheet.AddRow()
		for _, t := range record {
			row.AddCell().Value = t
		}
	}
	// 保存文件
	err := file.Save(path) // "file.xlsx"
	return err
}
