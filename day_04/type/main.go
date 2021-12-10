package main

import "fmt"

// 自定义类型和类型别名

type myInt int

type yourInt int

func main() {
	var a myInt
	a = 100
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	var b yourInt
	b = 200
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// 为什么需要类型别名：代码更清晰
	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
