package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

// 日志库开发
// 需求分析：支持往不同的地方输出日志
// 日志分级别：Debug、Trace、Info

// 判断是否输出日志方法，方法接收者为日志器
func (l Logger) enable(logLevel LogLevel) bool {
	return logLevel >= l.level
}

// Newlog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		level: level,
	}
}

func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, filePath, line := getInfo(3)
	fmt.Printf("[%s] [Debug] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, filePath, line, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		now := time.Now()
		fmt.Printf("[%s] [Trace] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [Info] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s] [Warning] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [Error] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s] [Fatal] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
