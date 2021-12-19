package main

import (
	"fmt"
	"time"
)

// 时间time包

// 时间戳
func timestampDemo() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func main() {
	now := time.Now() //获取当前时间对象
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	timestampDemo()
	// 使用time.Unix()可以把时间戳转换为时间
	// 时间格式化 go语言时间格式2006-01-02 15:04:05.000
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	// 按照格式将字符串时间转为时间戳
	timeObj, err := time.Parse("2006-01-02", "2000-02-18")
	if err != nil {
		fmt.Printf("parse error, err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	// sleep函数
	time.Sleep(5 * time.Second)
	fmt.Println("睡完了")
}
