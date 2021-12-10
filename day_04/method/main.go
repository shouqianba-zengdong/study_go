package main

import "fmt"

// 给自定义类型加方法
// 不能给别的包里面的类型定义方法

type myInt int

func (m myInt) hello() {
	fmt.Println("hello world")
}

func main() {
	m := myInt(100)
	m.hello()
}
