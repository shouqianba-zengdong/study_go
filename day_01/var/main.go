package main

import (
	"fmt"
)

// 主要方式一：全局批量声明
var (
	i int
	f float64
	b bool
	s string
)

func main() {
	// 变量声明：类型 + 值，可以只声明类型，变量初始值为零值
	i = 10
	f = 64.0
	b = true
	s = "Hello World"
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// 声明变量同时复制
	var hello string = "Hello World"
	fmt.Println(hello)
	// 声明变量只声明值，编译器自动判断类型
	var x = 1
	fmt.Println(x)
	// 主要方式二：简短变量声明，只能在函数中使用
	s3 := "哈哈哈"
	fmt.Println(s3)
	mydict := make(map[string]string)
	mydict["name"] = "曾东"
	for k := range mydict {
		fmt.Println(k)
	}
	// 除全局变量外，变量声明后必须使用
	// 一个作用域内，同一变量名只可声明一次
}
