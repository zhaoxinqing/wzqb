package controller

import (
	"Kilroy/app/constant"
	"encoding/csv"
	"encoding/json"
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
	// constant.ResSuccess(c, "postAdmin")
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
	_, err := os.Create("./docs/upload/")
	filePath := "./docs/upload/" + fmt.Sprintf("%d", time.Now().Unix()) + filename
	newFilePath := "./docs/upload/" + fmt.Sprintf("%d", time.Now().Unix()) + "m" + filename

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
	}

	// 写入新文件
	f, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	writer := csv.NewWriter(f)
	var header = []string{"grid_id", "city", "pred_sale_area", "score", "grid_type", "grid_size", "jingwei"}
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
		gerom = line[6]
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
			line[6] = str001

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
	constant.ResSuccess(c, "postAdmin")
}

func SortCSV(c *gin.Context) {
	var (
		// header = []string{"grid_id", "city", "province", "area", "pred_sale_area", "score", "grid_type", "grid_size", "jingwei", "prediction_explain"}
		path1  = "./docs/upload/" + "郑州.csv"
		path2  = "./docs/upload/" + "上海.csv"
		path3  = "./docs/upload/" + "深圳.csv"
		path4  = "./docs/upload/" + "武汉.csv"
		path5  = "./docs/upload/" + "北京.csv"
		path6  = "./docs/upload/" + "重庆.csv"
		path7  = "./docs/upload/" + "杭州.csv"
		path8  = "./docs/upload/" + "天津.csv"
		path9  = "./docs/upload/" + "合肥.csv"
		path10 = "./docs/upload/" + "青岛.csv"
		path11 = "./docs/upload/" + "成都.csv"
		path12 = "./docs/upload/" + "长沙.csv"
		path13 = "./docs/upload/" + "西安.csv"
		path14 = "./docs/upload/" + "宁波.csv"
		path15 = "./docs/upload/" + "苏州.csv"
		path16 = "./docs/upload/" + "南京.csv"
		path17 = "./docs/upload/" + "广州.csv"
		path18 = "./docs/upload/" + "佛山.csv"
		path19 = "./docs/upload/" + "东莞.csv"
		path20 = "./docs/upload/" + "厦门.csv"
	)

	file, err := c.FormFile("upload")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// err = os.Mkdir("docs/upload", 0666)
	uploadFilePath := "docs/upload/" + file.Filename

	err = c.SaveUploadedFile(file, uploadFilePath)

	// // 写入新文件
	// f1, _ := os.Create(path1)
	// f2, _ := os.Create(path2)
	// f3, _ := os.Create(path3)
	// f4, _ := os.Create(path4)
	// f5, _ := os.Create(path5)
	// f6, _ := os.Create(path6)
	// f7, _ := os.Create(path7)
	// f8, _ := os.Create(path8)
	// f9, _ := os.Create(path9)
	// f10, _ := os.Create(path10)
	// f11, _ := os.Create(path11)
	// f12, _ := os.Create(path12)
	// f13, _ := os.Create(path13)
	// f14, _ := os.Create(path14)
	// f15, _ := os.Create(path15)
	// f16, _ := os.Create(path16)
	// f17, _ := os.Create(path17)
	// f18, _ := os.Create(path18)
	// f19, _ := os.Create(path19)
	// f20, _ := os.Create(path20)

	// 追加
	f1, _ := os.OpenFile(path1, os.O_WRONLY|os.O_APPEND, 0666)
	f2, _ := os.OpenFile(path2, os.O_WRONLY|os.O_APPEND, 0666)
	f3, _ := os.OpenFile(path3, os.O_WRONLY|os.O_APPEND, 0666)
	f4, _ := os.OpenFile(path4, os.O_WRONLY|os.O_APPEND, 0666)
	f5, _ := os.OpenFile(path5, os.O_WRONLY|os.O_APPEND, 0666)
	f6, _ := os.OpenFile(path6, os.O_WRONLY|os.O_APPEND, 0666)
	f7, _ := os.OpenFile(path7, os.O_WRONLY|os.O_APPEND, 0666)
	f8, _ := os.OpenFile(path8, os.O_WRONLY|os.O_APPEND, 0666)
	f9, _ := os.OpenFile(path9, os.O_WRONLY|os.O_APPEND, 0666)
	f10, _ := os.OpenFile(path10, os.O_WRONLY|os.O_APPEND, 0666)
	f11, _ := os.OpenFile(path11, os.O_WRONLY|os.O_APPEND, 0666)
	f12, _ := os.OpenFile(path12, os.O_WRONLY|os.O_APPEND, 0666)
	f13, _ := os.OpenFile(path13, os.O_WRONLY|os.O_APPEND, 0666)
	f14, _ := os.OpenFile(path14, os.O_WRONLY|os.O_APPEND, 0666)
	f15, _ := os.OpenFile(path15, os.O_WRONLY|os.O_APPEND, 0666)
	f16, _ := os.OpenFile(path16, os.O_WRONLY|os.O_APPEND, 0666)
	f17, _ := os.OpenFile(path17, os.O_WRONLY|os.O_APPEND, 0666)
	f18, _ := os.OpenFile(path18, os.O_WRONLY|os.O_APPEND, 0666)
	f19, _ := os.OpenFile(path19, os.O_WRONLY|os.O_APPEND, 0666)
	f20, _ := os.OpenFile(path20, os.O_WRONLY|os.O_APPEND, 0666)

	writer1 := csv.NewWriter(f1)
	writer2 := csv.NewWriter(f2)
	writer3 := csv.NewWriter(f3)
	writer4 := csv.NewWriter(f4)
	writer5 := csv.NewWriter(f5)
	writer6 := csv.NewWriter(f6)
	writer7 := csv.NewWriter(f7)
	writer8 := csv.NewWriter(f8)
	writer9 := csv.NewWriter(f9)
	writer10 := csv.NewWriter(f10)
	writer11 := csv.NewWriter(f11)
	writer12 := csv.NewWriter(f12)
	writer13 := csv.NewWriter(f13)
	writer14 := csv.NewWriter(f14)
	writer15 := csv.NewWriter(f15)
	writer16 := csv.NewWriter(f16)
	writer17 := csv.NewWriter(f17)
	writer18 := csv.NewWriter(f18)
	writer19 := csv.NewWriter(f19)
	writer20 := csv.NewWriter(f20)

	// writer1.Write(header)
	// writer2.Write(header)
	// writer3.Write(header)
	// writer4.Write(header)
	// writer5.Write(header)
	// writer6.Write(header)
	// writer7.Write(header)
	// writer8.Write(header)
	// writer9.Write(header)
	// writer10.Write(header)
	// writer11.Write(header)
	// writer12.Write(header)
	// writer13.Write(header)
	// writer14.Write(header)
	// writer15.Write(header)
	// writer16.Write(header)
	// writer17.Write(header)
	// writer18.Write(header)
	// writer19.Write(header)
	// writer20.Write(header)

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
		gerom = line[8]
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
			line[8] = str001
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
			case "重庆市":
				writer6.Write(line)
				writer6.Flush()
			case "杭州市":
				writer7.Write(line)
				writer7.Flush()
			case "天津市":
				writer8.Write(line)
				writer8.Flush()
			case "合肥市":
				writer9.Write(line)
				writer9.Flush()
			case "青岛市":
				writer10.Write(line)
				writer10.Flush()
			case "成都市":
				writer11.Write(line)
				writer11.Flush()
			case "长沙市":
				writer12.Write(line)
				writer12.Flush()
			case "西安市":
				writer13.Write(line)
				writer13.Flush()
			case "宁波市":
				writer14.Write(line)
				writer14.Flush()
			case "苏州市":
				writer15.Write(line)
				writer15.Flush()
			case "南京市":
				writer16.Write(line)
				writer16.Flush()
			case "广州市":
				writer17.Write(line)
				writer17.Flush()
			case "佛山市":
				writer18.Write(line)
				writer18.Flush()
			case "东莞市":
				writer19.Write(line)
				writer19.Flush()
			case "厦门市":
				writer20.Write(line)
				writer20.Flush()
			}
		}
	}
	constant.ResSuccess(c, "Success")
}

func SortFeature(c *gin.Context) {
	file, _ := c.FormFile("upload")
	filename := file.Filename
	_, err := os.Create("./docs/upload/")
	filePath := "./docs/upload/" + fmt.Sprintf("%d", time.Now().Unix()) + filename
	newFilePath := "./docs/upload/" + fmt.Sprintf("%d", time.Now().Unix()) + "new" + filename

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
	}
	// 写入新文件
	f, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	writer := csv.NewWriter(f)
	var header = []string{"grid_id", "city", "pred_sale_area", "score", "grid_type", "grid_size", "prediction_explain", "Business", "Shopping", "People", "Traffic", "Home"}
	writer.Write(header)
	writer.Flush()
	//打开流
	clientsFile, err := os.Open(filePath)
	reader := csv.NewReader(clientsFile)
	for {
		var featureStr string
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		featureStr = line[9]
		if len(featureStr) > 30 {
			var features Features
			json.Unmarshal([]byte(featureStr), &features)
			fmt.Println(features)

			// 拆分数据
			var a Features02
			b := features
			// 办公
			a.Business = b.FAvgFloor1k + b.FAvgFloors1k + b.FMaxFloor1k + b.FMaxFloors1k + b.FMinFloor1k + b.FMinFloors1k + b.FTotalFloor1k + b.FTotalFloors1k +
				b.FAvgRent1k + b.FMaxRent1k + b.FMinRent1k + b.FCntOfficeBuilding1k + b.FPropertyTypes1k + b.FTotalSattledEnterprise1k + b.FBussinessLevel1k + b.FBussinessType1k
			line[10] = fmt.Sprint(a.Business)
			// 商业
			a.Shopping = b.FAvgCommercialarea1k + b.FMaxCommercialarea1k + b.FMinCommercialarea1k + b.FTotalCommercialarea1k + b.FCity + b.FShopCnt1k
			line[11] = fmt.Sprint(a.Shopping)
			// 人口
			a.People = b.FG400001_1k + b.FG400002_1k + b.FG400003_1k + b.FG400004_1k + b.FG400005_1k + b.FG400006_1k + b.FG400007_1k + b.FG400008_1k + b.FG400009_1k + b.FG400010_1k + b.FG400011_1k + b.FG400011_1k +
				b.FRtoG400001_1k + b.FRtoG400002_1k + b.FRtoG400003_1k + b.FRtoG400004_1k + b.FRtoG400005_1k + b.FRtoG400006_1k + b.FRtoG400007_1k + b.FRtoG400008_1k + b.FRtoG400009_1k + b.FRtoG400010_1k + b.FRtoG400011_1k + b.FRtoG400012_1k +
				b.FGFam1gHu1k + b.FGFam3gHu1k + b.FGFam2gHu1k + b.FGFam4gHu1k + b.FGFamHu1k + b.FRtoGFam1gHu1k + b.FRtoGFam2gHu1k + b.FRtogFam3gHu1k + b.FRtoGFam4gHu1k +
				b.FGA100001_1k + b.FGA200001_1k + b.FGA200002_1k + b.FGA200003_1k + b.FGA200004_1k + b.FGA200005_1k + b.FGA200007_1k +
				b.FGA200008_1k + b.FGA200009_1k + b.FGA200010_1k + b.FGA200011_1k + b.FGA200012_1k + b.FGA200013_1k + b.FGA200014_1k +
				b.FGA200015_1k + b.FGA200016_1k + b.FGA200017_1k + b.FGA200019_1k + b.FGA200020_1k + b.FGA200021_1k + b.FGA200023_1k +
				b.FGA200024_1k + b.FGA200025_1k + b.FGA200026_1k + b.FGA200027_1k + b.FGA200028_1k + b.FGA200029_1k + b.FGA200030_1k +
				b.FGA200031_1k + b.FGA200032_1k + b.FGA200033_1k + b.FGA200034_1k + b.FGA200035_1k + b.FGA200036_1k + b.FGA200037_1k + b.FGA200038_1k +
				b.FRtoGA200001_1k + b.FRtoGA200002_1k + b.FRtoGA200003_1k + b.FRtoGA200004_1k + b.FRtoGA200005_1k + b.FRtoGA200006_1k + b.FRtoGA200007_1k + b.FRtoGA200008_1k +
				b.FRtoGA200009_1k + b.FRtoGA200010_1k + b.FRtoGA200011_1k + b.FRtoGA200012_1k + b.FRtoGA200013_1k + b.FRtoGA200014_1k + b.FRtoGA200015_1k + b.FRtoGA200016_1k +
				b.FRtoGA200017_1k + b.FRtoGA200018_1k + b.FRtoGA200019_1k + b.FRtoGA200020_1k + b.FRtoGA200021_1k + b.FRtoGA200022_1k + b.FRtoGA200023_1k + b.FRtoGA200024_1k +
				b.FRtoGA200025_1k + b.FRtoGA200026_1k + b.FRtoGA200027_1k + b.FRtoGA200028_1k + b.FRtoGA200029_1k + b.FRtoGA200030_1k + b.FRtoGA200031_1k + b.FRtoGA200032_1k +
				b.FRtoGA200033_1k + b.FRtoGA200034_1k + b.FRtoGA200035_1k + b.FRtoGA200036_1k + b.FRtoGA200037_1k + b.FRtoGA200038_1k
			line[12] = fmt.Sprint(a.People)
			// 交通
			a.Traffic = b.FCntSubway1k + b.FCntBus1k
			line[13] = fmt.Sprint(a.Traffic)
			// 住宅be
			a.Home = b.FAvgPrice1k + b.FAvgAfforestRate1k + b.FAvgPlotRatio1k + b.FMaxPlotRatio1k + b.FMinPlotRatio1k +
				b.FAvgPropertyPrice1k + b.FMaxPropertyPrice1k + b.FMinPropertyPrice1k +
				b.FCntPlot1k + b.FAvgArea1k + b.FAvgDArea1k + b.FMaxArea1k + b.FMaxDArea1k + b.FMinArea1k + b.FMinDArea1k + b.FTotalArea1k
			line[14] = fmt.Sprint(a.Home)
		}
		writer.Flush()
	}
	constant.ResSuccess(c, "postAdmin")
}

type Features struct {
	// 办公
	FAvgFloor1k               float64 `json:"f_avg_floor_1k"`                //平均楼层
	FAvgFloors1k              float64 `json:"f_avg_floors_1k"`               // 楼层
	FMaxFloor1k               float64 `json:"f_max_floor_1k"`                // 最大楼层
	FMaxFloors1k              float64 `json:"f_max_floors_1k"`               // 最大楼层
	FMinFloor1k               float64 `json:"f_min_floor_1k"`                //楼层
	FMinFloors1k              float64 `json:"f_min_floors_1k"`               // 楼层
	FTotalFloor1k             float64 `json:"f_total_floor_1k"`              // 楼层
	FTotalFloors1k            float64 `json:"f_total_floors_1k"`             // 总楼层
	FAvgRent1k                float64 `json:"f_avg_rent_1k"`                 // 平均租金
	FMaxRent1k                float64 `json:"f_max_rent_1k"`                 // 最大
	FMinRent1k                float64 `json:"f_min_rent_1k"`                 // 最小
	FCntOfficeBuilding1k      float64 `json:"f_cnt_office_building_1k"`      // 写字楼
	FPropertyTypes1k          float64 `json:"f_property_types_1k"`           // 物业类型
	FTotalSattledEnterprise1k float64 `json:"f_total_sattled_enterprise_1k"` // 总入住企业
	FBussinessLevel1k         float64 `json:"f_bussiness_level_1k"`          // 业务水平
	FBussinessType1k          float64 `json:"f_bussiness_type_1k"`           // 业务类型

	// 商业
	FAvgCommercialarea1k   float64 `json:"f_avg_commercialarea_1k"`   // 平均商业面积
	FMaxCommercialarea1k   float64 `json:"f_max_commercialarea_1k"`   // 最大商业区
	FMinCommercialarea1k   float64 `json:"f_min_commercialarea_1k"`   // 最小商业区
	FTotalCommercialarea1k float64 `json:"f_total_commercialarea_1k"` // 总商业区
	FCity                  float64 `json:"f_city"`                    // 城市
	FShopCnt1k             float64 `json:"f_shop_cnt_1k"`             // 购物中心

	// 住宅
	FAvgPrice1k         float64 `json:"f_avg_price_1k"`         // 挂牌价格
	FAvgAfforestRate1k  float64 `json:"f_avg_afforest_rate_1k"` // 绿化
	FAvgPlotRatio1k     float64 `json:"f_avg_plot_ratio_1k"`    // 容积率
	FMaxPlotRatio1k     float64 `json:"f_max_plot_ratio_1k"`
	FMinPlotRatio1k     float64 `json:"f_min_plot_ratio_1k"`
	FAvgPropertyPrice1k float64 `json:"f_avg_property_price_1k"` // 物业价格
	FMaxPropertyPrice1k float64 `json:"f_max_property_price_1k"`
	FMinPropertyPrice1k float64 `json:"f_min_property_price_1k"`
	FCntPlot1k          float64 `json:"f_cnt_plot_1k"` // 小区个数
	FAvgArea1k          float64 `json:"f_avg_area_1k"` // 平均占地面积
	FAvgDArea1k         float64 `json:"f_avg_d_area_1k"`
	FMaxArea1k          float64 `json:"f_max_area_1k"`
	FMaxDArea1k         float64 `json:"f_max_d_area_1k"` // 最大面积
	FMinArea1k          float64 `json:"f_min_area_1k"`   // 最小面积
	FMinDArea1k         float64 `json:"f_min_d_area_1k"`
	FTotalArea1k        float64 `json:"f_total_area_1k"` // 总面积

	// 交通
	FCntSubway1k float64 `json:"f_cnt_subway_1k"` // 地铁
	FCntBus1k    float64 `json:"f_cnt_bus_1k"`    // 公交车

	// 人口
	FG400001_1k    float64 `json:"f_g_400001_1k"` // 教育结构
	FG400002_1k    float64 `json:"f_g_400002_1k"`
	FG400003_1k    float64 `json:"f_g_400003_1k"`
	FG400004_1k    float64 `json:"f_g_400004_1k"`
	FG400005_1k    float64 `json:"f_g_400005_1k"`
	FG400006_1k    float64 `json:"f_g_400006_1k"`
	FG400007_1k    float64 `json:"f_g_400007_1k"`
	FG400008_1k    float64 `json:"f_g_400008_1k"`
	FG400009_1k    float64 `json:"f_g_400009_1k"`
	FG400010_1k    float64 `json:"f_g_400010_1k"`
	FG400011_1k    float64 `json:"f_g_400011_1k"`
	FG400012_1k    float64 `json:"f_g_400012_1k"`
	FRtoG400001_1k float64 `json:"f_rto_g_400001_1k"` // 教育结构（占比）
	FRtoG400002_1k float64 `json:"f_rto_g_400002_1k"`
	FRtoG400003_1k float64 `json:"f_rto_g_400003_1k"`
	FRtoG400004_1k float64 `json:"f_rto_g_400004_1k"`
	FRtoG400005_1k float64 `json:"f_rto_g_400005_1k"`
	FRtoG400006_1k float64 `json:"f_rto_g_400006_1k"`
	FRtoG400007_1k float64 `json:"f_rto_g_400007_1k"`
	FRtoG400008_1k float64 `json:"f_rto_g_400008_1k"`
	FRtoG400009_1k float64 `json:"f_rto_g_400009_1k"`
	FRtoG400010_1k float64 `json:"f_rto_g_400010_1k"`
	FRtoG400011_1k float64 `json:"f_rto_g_400011_1k"`
	FRtoG400012_1k float64 `json:"f_rto_g_400012_1k"`
	FGFam1gHu1k    float64 `json:"f_g_fam1g_hu_1k"` // 家庭户
	FGFam2gHu1k    float64 `json:"f_g_fam2g_hu_1k"`
	FGFam3gHu1k    float64 `json:"f_g_fam3g_hu_1k"`
	FGFam4gHu1k    float64 `json:"f_g_fam4g_hu_1k"`
	FGFamHu1k      float64 `json:"f_g_fam_hu_1k"`
	FRtoGFam1gHu1k float64 `json:"f_rto_g_fam1g_hu_1k"` // 一代户				// 家庭户(占比)
	FRtoGFam2gHu1k float64 `json:"f_rto_g_fam2g_hu_1k"` // 二代户
	FRtoGFam4gHu1k float64 `json:"f_rto_g_fam4g_hu_1k"` // 四代户
	FRtogFam3gHu1k float64 `json:"f_rtog_fam3g_hu_1k"`  // 三代户
	FGA100001_1k   float64 `json:"f_g_a100001_1k"`      // 人口年龄结构
	FGA200001_1k   float64 `json:"f_g_a200001_1k"`
	FGA200002_1k   float64 `json:"f_g_a200002_1k"`
	FGA200003_1k   float64 `json:"f_g_a200003_1k"`
	FGA200004_1k   float64 `json:"f_g_a200004_1k"`
	FGA200005_1k   float64 `json:"f_g_a200005_1k"`
	FGA200007_1k   float64 `json:"f_g_a200007_1k"`
	FGA200008_1k   float64 `json:"f_g_a200008_1k"`
	FGA200009_1k   float64 `json:"f_g_a200009_1k"`
	FGA200010_1k   float64 `json:"f_g_a200010_1k"`
	FGA200011_1k   float64 `json:"f_g_a200011_1k"`
	FGA200012_1k   float64 `json:"f_g_a200012_1k"`
	FGA200013_1k   float64 `json:"f_g_a200013_1k"`
	FGA200014_1k   float64 `json:"f_g_a200014_1k"`
	FGA200015_1k   float64 `json:"f_g_a200015_1k"`
	FGA200016_1k   float64 `json:"f_g_a200016_1k"`
	FGA200017_1k   float64 `json:"f_g_a200017_1k"`
	FGA200019_1k   float64 `json:"f_g_a200019_1k"`
	FGA200020_1k   float64 `json:"f_g_a200020_1k"`
	FGA200021_1k   float64 `json:"f_g_a200021_1k"`
	FGA200023_1k   float64 `json:"f_g_a200023_1k"`
	FGA200024_1k   float64 `json:"f_g_a200024_1k"`
	FGA200025_1k   float64 `json:"f_g_a200025_1k"`
	FGA200026_1k   float64 `json:"f_g_a200026_1k"`
	FGA200027_1k   float64 `json:"f_g_a200027_1k"`
	FGA200028_1k   float64 `json:"f_g_a200028_1k"`
	FGA200029_1k   float64 `json:"f_g_a200029_1k"`
	FGA200030_1k   float64 `json:"f_g_a200030_1k"`
	FGA200031_1k   float64 `json:"f_g_a200031_1k"`
	FGA200032_1k   float64 `json:"f_g_a200032_1k"`
	FGA200033_1k   float64 `json:"f_g_a200033_1k"`
	FGA200034_1k   float64 `json:"f_g_a200034_1k"`
	FGA200035_1k   float64 `json:"f_g_a200035_1k"`
	FGA200036_1k   float64 `json:"f_g_a200036_1k"`
	FGA200037_1k   float64 `json:"f_g_a200037_1k"`
	FGA200038_1k   float64 `json:"f_g_a200038_1k"`

	FRtoGA200001_1k float64 `json:"f_rto_g_a200001_1k"` // 人口年龄结构（占比）
	FRtoGA200002_1k float64 `json:"f_rto_g_a200002_1k"`
	FRtoGA200003_1k float64 `json:"f_rto_g_a200003_1k"`
	FRtoGA200004_1k float64 `json:"f_rto_g_a200004_1k"`
	FRtoGA200005_1k float64 `json:"f_rto_g_a200005_1k"`
	FRtoGA200006_1k float64 `json:"f_rto_g_a200006_1k"`
	FRtoGA200007_1k float64 `json:"f_rto_g_a200007_1k"`
	FRtoGA200008_1k float64 `json:"f_rto_g_a200008_1k"`
	FRtoGA200009_1k float64 `json:"f_rto_g_a200009_1k"`
	FRtoGA200010_1k float64 `json:"f_rto_g_a200010_1k"`
	FRtoGA200011_1k float64 `json:"f_rto_g_a200011_1k"`
	FRtoGA200012_1k float64 `json:"f_rto_g_a200012_1k"`
	FRtoGA200013_1k float64 `json:"f_rto_g_a200013_1k"`
	FRtoGA200014_1k float64 `json:"f_rto_g_a200014_1k"`
	FRtoGA200015_1k float64 `json:"f_rto_g_a200015_1k"`
	FRtoGA200016_1k float64 `json:"f_rto_g_a200016_1k"`
	FRtoGA200017_1k float64 `json:"f_rto_g_a200017_1k"`
	FRtoGA200018_1k float64 `json:"f_rto_g_a200018_1k"`
	FRtoGA200019_1k float64 `json:"f_rto_g_a200019_1k"`
	FRtoGA200020_1k float64 `json:"f_rto_g_a200020_1k"`
	FRtoGA200021_1k float64 `json:"f_rto_g_a200021_1k"`
	FRtoGA200022_1k float64 `json:"f_rto_g_a200022_1k"`
	FRtoGA200023_1k float64 `json:"f_rto_g_a200023_1k"`
	FRtoGA200024_1k float64 `json:"f_rto_g_a200024_1k"`
	FRtoGA200025_1k float64 `json:"f_rto_g_a200025_1k"`
	FRtoGA200026_1k float64 `json:"f_rto_g_a200026_1k"`
	FRtoGA200027_1k float64 `json:"f_rto_g_a200027_1k"`
	FRtoGA200028_1k float64 `json:"f_rto_g_a200028_1k"`
	FRtoGA200029_1k float64 `json:"f_rto_g_a200029_1k"`
	FRtoGA200030_1k float64 `json:"f_rto_g_a200030_1k"`
	FRtoGA200031_1k float64 `json:"f_rto_g_a200031_1k"`
	FRtoGA200032_1k float64 `json:"f_rto_g_a200032_1k"`
	FRtoGA200033_1k float64 `json:"f_rto_g_a200033_1k"`
	FRtoGA200034_1k float64 `json:"f_rto_g_a200034_1k"`
	FRtoGA200035_1k float64 `json:"f_rto_g_a200035_1k"`
	FRtoGA200036_1k float64 `json:"f_rto_g_a200036_1k"`
	FRtoGA200037_1k float64 `json:"f_rto_g_a200037_1k"`
	FRtoGA200038_1k float64 `json:"f_rto_g_a200038_1k"`
}

type Features02 struct {
	People   float64 `json:"people"`
	Business float64 `json:"business"`
	Traffic  float64 `json:"traffic"`
	Shopping float64 `json:"shopping"`
	Home     float64 `json:"home"`
}
