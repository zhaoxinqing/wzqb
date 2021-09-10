package server

import (
	"fmt"

	"time"

	"github.com/robfig/cron/v3"
)

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
