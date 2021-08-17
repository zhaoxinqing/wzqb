package util

import (
	"encoding/json"
	"fmt"
	"strconv"

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

// interface 转 string
func Strval(value interface{}) string {

	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
