package controller

import (
	"Moonlight/app/common"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/geojson"
)

func GenerateAOI(c *gin.Context) {
	var (
		filePath    = "./docs/aoi.csv"
		newFilePath = "./docs/new_aoi.csv"
		header      = []string{"经纬度", "wkt", "geojson"}
	)
	newfile, err := os.Create(newFilePath) // 打开新文件
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(newfile)
	writer.Write(header)
	writer.Flush() // 写入标题行

	file, err := os.Open(filePath) //打开原文件
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err == io.EOF { // 是否读结束
			break
		}
		if err != nil { // 是否读报错
			fmt.Println(err)
			return
		}
		coordinate := line[0] // 坐标
		if len(coordinate) > 10 {
			lngLat := strings.Split(coordinate, ",")
			lng, _ := strconv.ParseFloat(lngLat[0], 64) // 经度 longitude -lng
			lat, _ := strconv.ParseFloat(lngLat[1], 64) // 纬度 latitude -lat
			lng0 := fmt.Sprint(lng - 0.0001)
			lng1 := fmt.Sprint(lng + 0.0001)
			lat0 := fmt.Sprint(lat - 0.0001)
			lat1 := fmt.Sprint(lat + 0.0001)
			wktStr := "POLYGON((" + lng0 + " " + lat0 + "," + lng0 + " " + lat1 + "," + lng1 + " " + lat1 + "," + lng1 + " " + lat0 + "," + lng0 + " " + lat0 + "))"
			geoGeometry, err := wkt.UnmarshalString(wktStr) // wkt-> geojson
			if err != nil {
				fmt.Println(err)
			}
			point := geojson.NewGeometry(geoGeometry)
			value := *point
			bytes, _ := value.MarshalJSON()
			line[1] = wktStr
			line[2] = fmt.Sprint(string(bytes))
			writer.Write(line) // 写入更新后记录
			writer.Flush()
			err = writer.Error()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	common.ResSuccess(c, "postAdmin")
}
