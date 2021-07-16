package ctrl

import (
	"Kilroy/app/common"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type MyJsonName struct {
	ModelOutputScore float64 `json:"model_output_score"`
	ShapValuesBySlot struct {
		ExpectedValue interface{} `json:"expected_value"`
		ShapValues    []struct {
			Label string  `json:"label"`
			Value float64 `json:"value"`
		} `json:"shap_values"`
	} `json:"shap_values_by_slot"`
}

func DocFeature(c *gin.Context) {
	file, _ := c.FormFile("upload")
	filename := file.Filename
	// _, err := os.Create("./docs/upload/")
	filePath := "./docs/upload/" + fmt.Sprintf("%d", time.Now().Unix()) + filename
	newFilePath := "./docs/upload/" + "new" + fmt.Sprintf("%d", time.Now().Unix()) + filename

	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
	}
	// 写入新文件
	f, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	writer := csv.NewWriter(f)
	var header = []string{"grid_id", "city", "area", "province", "score", "pred_sale_area", "grid_type", "grid_size", "jingwei", "prediction_explain", "Business", "Shopping", "People", "Traffic", "Home", "All"}
	writer.Write(header)
	writer.Flush()
	//打开流
	clientsFile, err := os.Open(filePath)
	reader := csv.NewReader(clientsFile)
	for {
		// var featureStr string
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		// 一、修改geomtry
		jingwei := line[8]
		if len(jingwei) > 15 {
			strs_arr := strings.Split(jingwei, `;`)
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

			// writer.Write(line)
			// writer.Flush()
			// 将缓存中的内容写入到文件里
			if err = writer.Error(); err != nil {
				fmt.Println(err)
			}
		}
		// 二、拆分因子
		var mod = make(map[string]float64)
		featureStr := line[9]
		if len(featureStr) > 30 {
			var myJsonName MyJsonName
			json.Unmarshal([]byte(featureStr), &myJsonName)
			lists := myJsonName.ShapValuesBySlot.ShapValues
			for _, list := range lists {
				mod[list.Label] = list.Value
			}

			// 办公
			business := mod["f_avg_floor"] + mod["f_avg_floors"] + mod["f_max_floor"] + mod["f_max_floors"] + mod["f_min_floor"] +
				mod["f_min_floors"] + mod["f_total_floor"] + mod["f_total_floors"] + mod["f_avg_rent"] + mod["f_max_rent"] +
				mod["f_min_rent"] + mod["f_cnt_office_building"] + mod["f_property_types"] + mod["f_total_sattled_enterprise"] +
				mod["f_bussiness_level"] + mod["f_bussiness_type"]
			line = append(line, fmt.Sprint(business))

			// 	// 商业
			shopping := mod["f_avg_commercialarea"] + mod["f_max_commercialarea"] + mod["f_min_commercialarea"] + mod["f_total_commercialarea"] +
				mod["f_city"] + mod["f_shop_cnt"]
			line = append(line, fmt.Sprint(shopping))

			// 人口
			people := mod["f_g_400001"] + mod["f_g_400002"] + mod["f_g_400003"] + mod["f_g_400004"] + mod["f_g_400005"] + mod["f_g_400006"] +
				mod["f_g_400007"] + mod["f_g_400008"] + mod["f_g_400009"] + mod["f_g_400010"] + mod["f_g_400011"] + mod["f_g_400012"] +
				mod["f_rto_g_400001"] + mod["f_rto_g_400002"] + mod["f_rto_g_400003"] + mod["f_rto_g_400004"] + mod["f_rto_g_400005"] +
				mod["f_rto_g_400006"] + mod["f_rto_g_400007"] + mod["f_rto_g_400008"] + mod["f_rto_g_400009"] + mod["f_rto_g_400010"] +
				mod["f_rto_g_400011"] + mod["f_rto_g_400012"] +
				mod["f_g_fam1g_hu"] + mod["f_g_fam2g_hu"] + mod["f_g_fam3g_hu"] + mod["f_g_fam4g_hu"] + mod["f_g_fam_hu"] +
				mod["f_rto_g_fam1g_hu"] + mod["f_rto_g_fam2g_hu"] + mod["f_rto_g_fam4g_hu"] + mod["f_rtog_fam3g_hu"] +
				mod["f_g_a100001"] + mod["f_g_a200001"] + mod["f_g_a200002"] + mod["f_g_a200003"] + mod["f_g_a200004"] + mod["f_g_a200005"] + mod["f_g_a200007"] +
				mod["f_g_a200008"] + mod["f_g_a200009"] + mod["f_g_a200010"] + mod["f_g_a200011"] + mod["f_g_a200012"] + mod["f_g_a200013"] + mod["f_g_a200014"] +
				mod["f_g_a200015"] + mod["f_g_a200016"] + mod["f_g_a200017"] + mod["f_g_a200019"] + mod["f_g_a200020"] + mod["f_g_a200021"] + mod["f_g_a200023"] +
				mod["f_g_a200024"] + mod["f_g_a200025"] + mod["f_g_a200026"] + mod["f_g_a200027"] + mod["f_g_a200028"] + mod["f_g_a200029"] + mod["f_g_a200030"] +
				mod["f_g_a200031"] + mod["f_g_a200032"] + mod["f_g_a200033"] + mod["f_g_a200034"] + mod["f_g_a200035"] + mod["f_g_a200036"] + mod["f_g_a200037"] +
				mod["f_g_a200038"] + mod["f_rto_g_a200001"] + mod["f_rto_g_a200002"] + mod["f_rto_g_a200003"] + mod["f_rto_g_a200004"] + mod["f_rto_g_a200005"] +
				mod["f_rto_g_a200006"] + mod["f_rto_g_a200007"] + mod["f_rto_g_a200008"] + mod["f_rto_g_a200009"] + mod["f_rto_g_a200010"] + mod["f_rto_g_a200011"] +
				mod["f_rto_g_a200012"] + mod["f_rto_g_a200013"] + mod["f_rto_g_a200014"] +
				mod["f_rto_g_a200015"] + mod["f_rto_g_a200016"] +
				mod["f_rto_g_a200017"] + mod["f_rto_g_a200018"] + mod["f_rto_g_a200019"] + mod["f_rto_g_a200020"] + mod["f_rto_g_a200021"] + mod["f_rto_g_a200022"] +
				mod["f_rto_g_a200023"] + mod["f_rto_g_a200024"] +
				mod["f_rto_g_a200025"] + mod["f_rto_g_a200026"] + mod["f_rto_g_a200027"] + mod["f_rto_g_a200028"] + mod["f_rto_g_a200029"] + mod["f_rto_g_a200030"] +
				mod["f_rto_g_a200031"] + mod["f_rto_g_a200032"] +
				mod["f_rto_g_a200033"] + mod["f_rto_g_a200034"] + mod["f_rto_g_a200035"] + mod["f_rto_g_a200036"] + mod["f_rto_g_a200037"] + mod["f_rto_g_a200038"]
			line = append(line, fmt.Sprint(people))

			// 交通
			traffic := mod["f_cnt_subway"] + mod["f_cnt_bus"]
			line = append(line, fmt.Sprint(traffic))

			// 住宅be
			home := mod["f_avg_price"] + mod["f_avg_afforest_rate"] + mod["f_avg_plot_ratio"] + mod["f_max_plot_ratio"] +
				mod["f_min_plot_ratio"] + mod["f_avg_property_price"] + mod["f_max_property_price"] + mod["f_min_property_price"] +
				mod["f_cnt_plot"] + mod["f_avg_area"] + mod["f_avg_d_area"] + mod["f_max_area"] + mod["f_max_d_area"] +
				mod["f_min_area"] + mod["f_min_d_area"] + mod["f_total_area"]
			line = append(line, fmt.Sprint(home))
			all := business + shopping + people + traffic + home
			line = append(line, fmt.Sprint(all))
			writer.Write(line)
			writer.Flush()
		}
	}
	common.ResSuccess(c, "postAdmin")
}
