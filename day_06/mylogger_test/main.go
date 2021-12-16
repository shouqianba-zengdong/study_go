package main

import (
	"time"

	"wosai-inc.com/study_go/day_06/mylogger"
)

// 测试mylogger包
func main() {
	log := mylogger.NewLog("error")
	for {
		log.Debug("这是一条Debug级别的日志")
		log.Info("这是一条Info级别的日志")
		log.Warning("这是一条Warning级别的日志")
		log.Error("这是一条Error级别的日志")
		log.Fatal("这是一条Fatal级别的日志")
		time.Sleep(3 * time.Second)
	}
}
