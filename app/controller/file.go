package controller

import (
	"Kilroy/app/common"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
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

type client struct {
	GridID       string `json:"grid_id"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Area         string `json:"area"`
	JingWei      string `json:"jingwei"`
	PredSaleArea string `json:"pred_sale_area"`
	Score        string `json:"score"`
	GridType     string `json:"grid_type"`
	GridSize     string `json:"grid_size"`
}

func UploadCSV(c *gin.Context) {
	var (
	// clients []client
	)
	file, _ := c.FormFile("upload")
	filename := file.Filename
	filePath := "./docs/upload/" + fmt.Sprintf("%d", time.Now().Unix()) + filename
	filePath2 := "./docs/upload/" + fmt.Sprintf("%d", time.Now().Unix()) + "m" + filename
	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
	}

	// 写入新文件
	f, err := os.Create(filePath2)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	writer := csv.NewWriter(f)
	var header = []string{"grid_id", "city", "province", "area", "geometry", "pred_sale_area", "score", "grid_type", "grid_size", "prediction_explain"}
	writer.Write(header)
	//打开流
	clientsFile, err := os.Open(filePath)
	reader := csv.NewReader(clientsFile)
	for {
		var gerom string
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		// MULTIPOLYGON(((119.191542517377 29.9370940653411,119.193205716491 29.937091073857,119.193205571661 29.9354243638227,119.191542372492 29.9354273553102,119.191542517377 29.9370940653411)))
		// 113.146029,34.421098;113.155038,34.421098;113.155038,34.412089;113.146029,34.412089;113.146029,34.421098
		gerom = line[4]
		if len(gerom) > 15 {
			strs_arr := strings.Split(gerom, `;`)
			str001 := "MULTIPOLYGON((("
			for idx, strs := range strs_arr {
				str := strings.Split(strs, `,`)
				str001 = str001 + str[0] + " " + str[1]
				if idx != 4 {
					str001 += ","
				}
			}
			str001 += ")))"
			line[4] = str001

			writer.Write(line)
			writer.Flush()
			// 将缓存中的内容写入到文件里
			if err = writer.Error(); err != nil {
				fmt.Println(err)
			}
		}
		// line = append(line, "geometry")
	}

	//遍历clients，每个结构体参数用client来获取，并按照需求进行处理

	c.String(200, "Success")
	common.ResSuccess(c, "postAdmin")
}

func SortCSV(c *gin.Context) {
	var (
		header = []string{"grid_id", "city", "province", "area", "geometry", "pred_sale_area", "score", "grid_type", "grid_size", "prediction_explain"}
		path1  = "./docs/upload/" + "郑州.csv"
		path2  = "./docs/upload/" + "上海.csv"
		path3  = "./docs/upload/" + "深圳.csv"
		path4  = "./docs/upload/" + "武汉.csv"
		path5  = "./docs/upload/" + "北京.csv"
	)

	file, _ := c.FormFile("upload")
	uploadFilePath := "./docs/upload/" + file.Filename

	_ = c.SaveUploadedFile(file, uploadFilePath)

	// 写入新文件
	f1, _ := os.Create(path1)
	f2, _ := os.Create(path2)
	f3, _ := os.Create(path3)
	f4, _ := os.Create(path4)
	f5, _ := os.Create(path5)

	// 追加
	// f1, _ := os.OpenFile(path1, os.O_WRONLY|os.O_APPEND, 0666)
	// f2, _ := os.OpenFile(path2, os.O_WRONLY|os.O_APPEND, 0666)
	// f3, _ := os.OpenFile(path3, os.O_WRONLY|os.O_APPEND, 0666)
	// f4, _ := os.OpenFile(path4, os.O_WRONLY|os.O_APPEND, 0666)
	// f5, _ := os.OpenFile(path5, os.O_WRONLY|os.O_APPEND, 0666)

	writer1 := csv.NewWriter(f1)
	writer2 := csv.NewWriter(f2)
	writer3 := csv.NewWriter(f3)
	writer4 := csv.NewWriter(f4)
	writer5 := csv.NewWriter(f5)

	writer1.Write(header)
	writer2.Write(header)
	writer3.Write(header)
	writer4.Write(header)
	writer5.Write(header)

	//打开流
	clientsFile, _ := os.Open(uploadFilePath)
	reader := csv.NewReader(clientsFile)
	for {
		var gerom string
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		gerom = line[4]
		if len(gerom) > 15 {
			strs_arr := strings.Split(gerom, `;`)
			str001 := "MULTIPOLYGON((("
			for idx, strs := range strs_arr {
				str := strings.Split(strs, `,`)
				str001 = str001 + str[0] + " " + str[1]
				if idx != 4 {
					str001 += ","
				}
			}
			str001 += ")))"
			line[4] = str001
			switch line[1] {
			case "郑州市":
				writer1.Write(line)
				writer1.Flush()
			case "上海市":
				writer2.Write(line)
				writer2.Flush()
			case "深圳市":
				writer3.Write(line)
				writer3.Flush()
			case "武汉市":
				writer4.Write(line)
				writer4.Flush()
			case "北京市":
				writer5.Write(line)
				writer5.Flush()
			}
		}
	}
	common.ResSuccess(c, "Success")
}
