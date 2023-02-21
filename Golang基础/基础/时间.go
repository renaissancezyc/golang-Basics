package main

import (
	f "fmt"
	"time"
)

// time包为我们提供了一个数据类型time.Time(作为值使用)以及显示和测量时间和日期的功能函数。

func main() {

	t := time.Now()

	// 获取当前时间   	 		 2023-01-10 20:53:34.2771301 +0800 CST m=+0.010439901
	f.Println(t)

	// 获取现在的分钟   		 53
	f.Println(t.Minute())

	// UTC世界通用时间  		 2023-01-10 12:53:34.3225321 +0000 UTC
	f.Println(time.Now().UTC())
	f.Println(time.Time())

	// 获取年   				 2023
	f.Println(t.Year())

	// 获取月   				 January
	f.Println(t.Month())

	// 获取日  				     10
	f.Println(t.Day())

	// 格式化日期字符串   		 2023.10.01
	f.Printf("%4d.%02d.%02d\n", t.Year(), t.Day(), t.Month())

	// 如果需要在应用程序在经过一定时间或周期执行某项任务（事件处理的特例），则可以使用 time.After 或者 time.Ticker

}
