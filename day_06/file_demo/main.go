package main

import (
	"fmt"
	"os"
)

// 1.文件对象的类型
// 2.获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open fail failed, err:%s\n", err)
		return
	}
	// 获取文件对象的类型
	fmt.Printf("%T\n", fileObj)
	// 获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%s\n", err)
		return
	}
	fmt.Printf("文件的大小是：%dB\n", fileInfo.Size())
}
