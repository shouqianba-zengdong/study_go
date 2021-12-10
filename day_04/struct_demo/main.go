package main

import "fmt"

// 结构体
// 结构体是值类型的
// go语言传参永远是值传递
type People struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 定义变量
	var people People
	// 通过字段赋值
	people.name = "曾东"
	people.age = 18
	people.gender = "男"
	people.hobby = []string{"篮球", "足球"}
	fmt.Println(people)
	// 匿名结构体
	var s struct {
		name string
		age  int
	}
	s.age = 18
	s.name = "haha"
	fmt.Println(s)
}
