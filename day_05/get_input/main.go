package main

import (
	"bufio"
	"fmt"
	"os"
)

// 从标准输入获取用户输入

func useScan() {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是：%s\n", s)
}

func useBufio() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func main() {
	useScan()
}
