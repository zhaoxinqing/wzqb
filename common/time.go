package common

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
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

// string转time格式
func StringToFormatTime(timeStr string) *Time {
	commonTime, _ := time.Parse("2006-01-02", timeStr)
	v := Time(commonTime)
	return &v
}

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

//定时任务
func TimeToNotification() {
	Test()
}

//cron表达式，每秒一次,  秒 分 时 日 月 周
func Test() {
	c := cron.New(cron.WithSeconds())
	spec := "*/10 * * * * ?"
	c.AddFunc(spec, func() {
		fmt.Println(time.Now())
	})
	c.Start()
}
