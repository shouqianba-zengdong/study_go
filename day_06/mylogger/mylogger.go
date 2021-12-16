package mylogger

import (
	"fmt"
	"path"
	"runtime"
)

type LogLevel uint16

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
	UNKNOWN
)

// 日志结构体
type Logger struct {
	level LogLevel
}

func getInfo(skip int) (funcName, filePath string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	filePath = path.Base(file)
	return
}
