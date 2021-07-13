package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		fmt.Println("任务名称不能为空")
	}
	switch args[1] {
	case "test":
		fmt.Println("这是测试")
	case "del_temp_dir":
		delTempDir()
	default:
		fmt.Println("Do Nothing")
	}
	os.Exit(0)
}

// 清理临时目录
func delTempDir() {
	err := os.RemoveAll("./doc/temp")
	if err != nil {
		fmt.Println(err)
	}
}

// go run script/script.go aa
// go run script/script.go deleteReport
