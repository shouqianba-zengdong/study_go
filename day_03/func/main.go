package main

import "fmt"

// 函数
// 函数的意义：减少代码重复开发
func sum(x, y int) int {
	return x + y
}

// 有参数，没有返回值
func sum1(x, y int) {
	fmt.Println(x + y)
}

// 有返回值，无参数
func sum2() int {
	return 2
}

// 无参数，无返回值
func f() {
	fmt.Println("我是一个没有参数也没有返回值的函数")
}

// 可变长参数
func f1(x string, y ...int) {
	fmt.Println(y)
}

// 多个返回值
func sum3(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

func main() {
	// 注意：在一个命名函数中不能再声明一个命名函数
	r1 := sum(1, 2)
	fmt.Println(r1)
	r2 := sum2()
	fmt.Println(r2)
	sum1(1, 2)
	f()
	f1("haha", 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sum3(2, 1))
}
