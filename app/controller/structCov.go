package controller

import (
	"Moonlight/app/common"
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/copier"
)

type TestGetScoreParam_2 struct {
	A_5  string `json:"a_5"`
	A_10 string `json:"a_10"`
	B_5  string `json:"b_100"`
	B_10 string `json:"b_50"`
}

type TestGetScoreParam struct {
	A_5  int `json:"a_5"`
	A_10 int `json:"a_10"`
	B_5  int `json:"b_100"`
	B_10 int `json:"b_50"`
}

type AMD_5 struct {
	A5 int `json:"a"  copier:"A_5"`
	B5 int `json:"b"  copier:"B_5"`
}

type AMD_10 struct {
	A10 int `json:"a" copier:"A_10"`
	B10 int `json:"b" copier:"B_10"`
}

// GetByTime 通过时间获取
func StructToStruct(c *gin.Context) {
	var (
		result = TestGetScoreParam{}
		body   = TestGetScoreParam_2{}
	)
	a := AMD_5{
		A5: 4,
		B5: 4,
	}
	b := AMD_10{
		A10: 5789,
		B10: 5789,
	}
	copier.Copy(&result, &a)
	copier.Copy(&result, &b)

	if err := gconv.Struct(result, &body); err != nil {
		fmt.Println(err)
	}
}

func StrTime(c *gin.Context) {
	//  获取id参数
	time, _ := c.GetQuery("time")
	var a string
	if len(time) > 20 {
		a = time[0:10] + " " + time[11:19]
	}

	fmt.Println(a)
	n := 14.6455
	num := math.Pow(2, n)
	common.ResSuccess(c, num)
}
