package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 自己定义读多少内容
func readFromFile() {
	// 打开文件
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed , err : %v\n", err)
		return
	}
	// 提前注册关闭文件
	defer fileObj.Close()
	// 定义字节切片存储读的内容
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		// 判断文件是否读完，是则return
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		// 判断文件读取是否出错，出错则给出提示并return
		if err != nil {
			fmt.Printf("read file failed, err : %v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		// 把读取的字节打印出来
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// 利用bufio每次读一行
func readFromBufio() {
	// 获取文件对象
	fileObj, err := os.Open("./main.go")
	// 判断获取文件对象是否出错，有错则提示并return
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 提前注册关闭文件对象行为
	defer fileObj.Close()
	// 获取读取器对象
	reader := bufio.NewReader(fileObj)
	for {
		// 换行符
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromFile()
	// readFromBufio()
	readFromFileByIoutil()
}
