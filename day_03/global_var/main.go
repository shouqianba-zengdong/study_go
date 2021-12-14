package main

import "fmt"

// 变量作用域

var x = 100 //定义一个全局变量

func f1() {
	// 函数中查找变量的顺序
	// 1.先在函数内部查找
	// 2.找不到就往外层函数查找，直到全局变量
	fmt.Println(x)
	// 函数内定义的变量，只能在函数内使用
	name := "曾东"
	fmt.Println(name)
}

func main() {
	f1()
	// 语句块作用域
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
