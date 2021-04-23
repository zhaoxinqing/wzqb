package main

import (
	"fmt"
	"os"
)

func DeleteReportFiles() {
	//删除目录
	err := os.RemoveAll("./docs/report")
	if err != nil {
		fmt.Println(err)
	}
}
