package utils

import (
	"fmt"
	"time"
)

func GetTime(msg string) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println(msg, time.Since(start))
	}(start)
}
