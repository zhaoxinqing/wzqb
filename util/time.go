package util

import (
	"fmt"
	"time"
)

// 时间转换
type Time time.Time

const timeFormart = "2006-01-02 15:04:05"

// 转换成时间戳
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// 转换成自定义格式
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func GetTime(msg string) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println(msg, time.Since(start))
	}(start)
}
