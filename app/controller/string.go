package controller

import (
	"fmt"
	"strings"
)

func GetString() {
	fmt.Println("嗨客网(www.haicoder.net)")
	// https://haicoder.net/golang/golang-string-split.html

	// 计数
	//使用 Strings.count() 函数，统计字符串中单个字符出现的次数
	strHaiCoder01 := "Study Golang From HaiCoder"
	count := strings.Count(strHaiCoder01, "o")
	fmt.Println("count =", count)

	// 分割
	//使用 strings.Fields 函数，实现按空格分割字符串
	strHaiCoder02 := "嗨客网 Hello HaiCoder"
	strArr := strings.Fields(strHaiCoder02)
	fmt.Println("strArr =", strArr)

	//使用 strings.Split 函数，实现按字符串分割字符串
	strHaiCoder := "Hello,HaiCoder Hello,World"
	strArr03 := strings.Split(strHaiCoder, "Hello")
	fmt.Println("strArr =", strArr03)

}
