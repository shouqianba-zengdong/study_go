package main

import "fmt"

// 结构体匿名字段

type Student struct {
	int
	string
}

func main() {
	student := Student{
		18,
		"曾东",
	}
	fmt.Println(student)
	// 通过类型访问字段
	// 所以一个结构体中每种类型的匿名字段只能有一个
	fmt.Println(student.int)
	fmt.Println(student.string)
}
