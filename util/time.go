package util

import (
	"fmt"
	"time"
)

// 时间转换
type Time time.Time

const TimeFormart = "2006-01-02 15:04:05"

func GetBetweenTime(msg string) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println(msg, time.Since(start))
	}(start)
}
