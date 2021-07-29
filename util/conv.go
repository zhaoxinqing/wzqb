package util

import (
	"fmt"

	"github.com/gogf/gf/util/gconv"
)

type Con_1 struct {
	Name   string
	Age    int64
	Sex    int64
	Remark string
}

type Con_2 struct {
	Name   string
	Age    int64
	Sex    int64
	Remark string
}

func ConStruct() {
	var (
		con_1 Con_1
		con_2 Con_2
	)
	con_1 = Con_1{
		Name:   "哈哈",
		Age:    18,
		Sex:    0,
		Remark: "哈哈的备注",
	}
	err = gconv.Struct(con_1, &con_2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(con_2)
}
