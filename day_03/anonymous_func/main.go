package main

import "fmt"

// 匿名函数
// 匿名函数一般用于函数内部

func main() {
	f1 := func(x, y int) int {
		return x + y
	}
	fmt.Println(f1(10, 20))
	// 匿名函数写完直接调用，不用变量保存，只使用一次
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("我是匿名函数，写完可以直接调用！")
	}(100, 200)
}
