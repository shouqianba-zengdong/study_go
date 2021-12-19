package main

import (
	"time"

	"wosai-inc.com/study_go/day_06/mylogger"
)

// 测试mylogger包
func main() {
	// 创建一个日志器：调用日志器构造函数，传入日志器级别
	id := 10000
	name := "曾东"
	// consoleLog := mylogger.NewConsoleLogger("info")
	log := mylogger.NewFileLogger("info", "./", "cengdong.log", 10*1024)
	for {
		// 调用log日志器的DEBUG方法
		log.Debug("这是一条Debug级别的日志")
		// 调用log日志器的Info方法
		log.Info("这是一条Info级别的日志")
		// 调用log日志器的Warning方法
		log.Warning("这是一条Warning级别的日志")
		// 调用log日志器的Error方法
		log.Error("这是一条Error级别的日志id:%d name:%s", id, name)
		// 调用log日志器的Fatal方法
		log.Fatal("这是一条Fatal级别的日志")
		time.Sleep(3 * time.Second)
	}
}
