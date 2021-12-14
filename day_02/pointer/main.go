package main

import (
	"fmt"
)

// 有关指针的操作有两种：&（取地址）*（根据地址取值）

func main() {
	n := 10
	p := &n
	fmt.Printf("%T\n", p)
	fmt.Println(*p)
	// 定义int类型指针，没有初始化，不可直接赋值
	var a *int
	fmt.Println(a)
	// new函数申请内存地址
	var b = new(int)
	fmt.Println(b)
	*b = 100
	fmt.Println(*b)
	// make函数也用于内存分配，但是只用于map、slice、chan类型
}
