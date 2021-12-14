package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("我是一个没有参数也没有返回值的函数")
}
func f2() (x int) {
	return 5
}

// 函数可以作为参数

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
	fmt.Println("我是一个参数为函数的函数")
}

func f4(x, y int) int {
	return x + y
}

// 函数可以作为返回值
func f5(x func() int) (ff func(int, int) int) {
	return ff
}

func main() {
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)
	f3(f2)
	fmt.Printf("%T\n", f5(f2))
}
