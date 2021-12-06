package main

import (
	"fmt"
)

func main() {
	f1 := 1.23456
	// go语言默认64位浮点型
	fmt.Printf("%T\n", f1)
	// 显示声明32位浮点型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
}
