package excel

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	Xlsx = "xlsx"
	Csv  = "csv"
)

func ExportTable(records [][]string, fileName string, fileType string) (savePath string, err error) {
	_ = os.MkdirAll("./doc/report/", os.ModePerm)
	var (
		prefixPath = "./doc/report/" + fmt.Sprintf("%s%d", fileName, time.Now().Unix())
	)
	switch fileType {
	case "xlsx":
		savePath = prefixPath + ".xlsx"
		err = SaveExcelFile(savePath, records)
	case "csv":
		savePath = prefixPath + ".csv"
		err = SaveCsvFile(savePath, records)
	}
	return
}

// SaveExcelFile ...
func SaveExcelFile(savePath string, records [][]string) (err error) {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	for i, record := range records {
		f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i+1), &record)
	}
	f.SetActiveSheet(index)
	err = f.SaveAs(savePath)

	return err
}

// SaveCsvFile ...
func SaveCsvFile(savePath string, records [][]string) (err error) {
	// 不存在则创建;存在则清空;读写模式;
	file, err := os.Create(savePath)
	if err != nil {
		return
	}
	// 延迟关闭
	defer file.Close()

	// 写入UTF-8 BOM，防止中文乱码
	_, _ = file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)

	// 写入数据
	for _, row := range records {
		_ = w.Write(row)
	}
	w.Flush()
	return err
}
