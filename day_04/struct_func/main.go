package main

import "fmt"

// 构造函数：返回一个结构体

type People struct {
	name string
	age  int
}

func newPeople(name string, age int) *People {
	return &People{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPeople("曾东", 18)
	p2 := newPeople("曾西", 18)
	fmt.Println(p1, p2)
}
