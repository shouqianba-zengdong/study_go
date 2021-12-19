package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里写日志

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// fileLogger构造函数
func NewFileLogger(levelStr, filePath, fileName string, maxFileSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       logLevel,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	f1.initFile()
	return f1
}

func (f *FileLogger) initFile() error {
	fullFIleName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFIleName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%s\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFIleName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%s\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) CloseFile() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%s\n", err)
		return false
	}
	// 如果当前文件的大小 大于等于 日志器设置的最大值，则返回true
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, filePath, line := getInfo(3)
		if f.checkSize(f.fileObj) {
			// f.splitFile()
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, filePath, line, msg)
		if lv >= ERROR {
			// 如果要记录的日志大于等于ERROR级别，我还要在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, filePath, line, msg)
		}
	}
}

// func (f *FileLogger) splitFile(file *os.File) (*os.File, err) {
// 	// 需要切割日志文件
// 	// 1.关闭当前的文件日志
// 	file.Close()
// 	// 2.备份一下 rename
// 	nowStr := time.Now().Format("20060102150405000")
// 	logName := path.Join(f.filePath, f.fileName)
// 	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
// 	os.Rename(logName, newLogName)
// 	// 3.打开一个新的日志文件
// 	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Printf("open file failed, err:%s\n", err)
// 		return nil, err
// 	}
// 	f.fileObj = fileObj
// 	// 4.将打开的日志文件对象赋值给 f.fileObj
// }

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
