package common

import (
	"strconv"
)

// StrToFloat64 ...
func StrToFloat64(str string) (floatValue float64, err error) {
	floatValue, err = strconv.ParseFloat(str, 64)
	return
}

// StrToInt ...
func StrToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

// StrToInt64 ...
func StrToInt64(str string) (int64, error) {
	intNum, err := strconv.Atoi(str)
	int64Num := int64(intNum)
	return int64Num, err
}
