package server

import (
	"fmt"

	"time"

	"github.com/robfig/cron/v3"
)

//定时任务
func TimeToNotification() {
	Test()
	TimeTask01()
	TimeTask02()
	TimeTask03()
}

//cron表达式，每秒一次,  秒 分 时 日 月 周
func Test() {
	c := cron.New(cron.WithSeconds())
	spec := "*/10 * * * * ?"
	c.AddFunc(spec, func() {
		fmt.Println(time.Now())
		test01()
	})
	c.Start()
}

//
func TimeTask01() {
	c := cron.New(cron.WithSeconds())
	spec := "*/* * * 5 * ?"
	c.AddFunc(spec, func() {
		fmt.Println("定时通知任务已启动")
	})
	c.Start()
}

//
func TimeTask02() {
	c := cron.New(cron.WithSeconds())
	spec := "*/* * * 60 * ?"
	c.AddFunc(spec, func() {
		fmt.Println("定时通知任务已启动")
	})
	c.Start()
}

//
func TimeTask03() {
	c := cron.New(cron.WithSeconds())
	spec := "*/* * * 30 * ?"
	c.AddFunc(spec, func() {
		fmt.Println("定时通知任务已启动")
	})
	c.Start()
}

func test01() {

}
