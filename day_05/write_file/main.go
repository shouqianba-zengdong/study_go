package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 向文件中写

func writeDemo1() {
	// os.OpenFile()三个参数，目标文件路径，
	fileObj, err := os.OpenFile("xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	str := "hello world"
	fileObj.Write([]byte(str))         // 写入字节切片数据
	fileObj.WriteString("hello world") //直接写入字符串数据
}

func writeDemo2() {
	// bufio.newWrite()
	fileObj, err := os.OpenFile("xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\f", err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello曾东\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

func writeDemo3() {
	str := "hello world"
	err := ioutil.WriteFile("./xxx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err :", err)
		return
	}
}

func main() {
	// writeDemo1()
	// writeDemo2()
	writeDemo3()
}
