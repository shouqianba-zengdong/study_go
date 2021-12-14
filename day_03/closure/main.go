package main

import "fmt"

// 闭包 = 函数 + 外部变量引用
// 闭包的底层原理
// 1.函数可以作为返回值
// 2.函数内部查找变量的方式

func adder(x int) func(y int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder(100)
	ret2 := ret(200)
	fmt.Println(ret2)
}
