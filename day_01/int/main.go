package main

import (
	"fmt"
)

func main() {
	i1 := 101
	// 十进制
	fmt.Printf("%d\n", i1)
	// 二进制
	fmt.Printf("%b\n", i1)
	// 八进制
	fmt.Printf("%x\n", i1)
	// 十六进制
	fmt.Printf("%o\n", i1)
	// 定义八进制数
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 定义十六进制数
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
}
