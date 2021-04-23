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
		fmt.Errorf("任务名称不能为空")
	} else {
		switch args[1] {
		case "aa":
			fmt.Println("这是测试")
		case "deleteReport":
			DeleteReportFiles()
		default:
			fmt.Errorf("Do Nothing")
		}
	}
	os.Exit(0)
}





