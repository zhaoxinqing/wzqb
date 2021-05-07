package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		args = os.Args
	)
	if len(args) < 1 {
		fmt.Println("任务名称不能为空")
	} else {
		switch args[1] {
		case "aa":
			fmt.Println("这是测试")
		case "deleteReport":
			deleteReportFiles()
		default:
			fmt.Println("Do Nothing")
		}
	}
	os.Exit(0)
}

// go run script/script.go aa
// go run script/script.go deleteReport

//删除目录
func deleteReportFiles() {
	err := os.RemoveAll("./docs/report")
	if err != nil {
		fmt.Println(err)
	}
}
