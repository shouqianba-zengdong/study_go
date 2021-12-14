package main

import (
	"fmt"
)

func main() {
	n := 10
	// 查看类型
	fmt.Printf("%T\n", n)
	// 默认值
	fmt.Printf("%v\n", n)
	// 二进制
	fmt.Printf("%b\n", n)
	// 十进制
	fmt.Printf("%d\n", n)
	// 八进制
	fmt.Printf("%o\n", n)
	// 十六进制
	fmt.Printf("%x\n", n)
	s := "Hello World"
	// 字符串
	fmt.Printf("%s\n", s)
	// 默认值
	fmt.Printf("%v\n", s)
}
