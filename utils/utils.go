package utils

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

// time
func Util(a string) string {
	// time
	timeNow := time.Now().Format(TimeFormart)
	fmt.Println(timeNow)
	return timeNow

}
