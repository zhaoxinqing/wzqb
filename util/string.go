package util

import (
	"fmt"
	"strconv"
	"strings"
)

// StrToFloat64 ...
func StrToFloat64(str string) (floatValue float64, err error) {
	floatValue, err = strconv.ParseFloat(str, 64)
	return
}

// StrToInt ...
func StrToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// StrToInt64 ...
func StrToInt64(str string) (int64, error) {
	intNum, err := strconv.Atoi(str)
	int64Num := int64(intNum)
	return int64Num, err
}

// StrToUint64 ...
func StrToUint64(str string) (uint64, error) {
	intNum, err := strconv.Atoi(str)
	uint64Num := uint64(intNum)
	return uint64Num, err
}

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
