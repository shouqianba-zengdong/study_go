package main

import "fmt"

// go语言中函数是拷贝

type People struct {
	name, gender string
	age          int
}

func f1(p People) {
	p.gender = "女"
	fmt.Printf("看看我在这个函数中被改成什么性别了:%s\n", p.gender)
}

func f2(p *People) {
	p.gender = "女"
}

func main() {
	var people People
	people.name = "曾东"
	people.gender = "男"
	people.age = 18
	f1(people)
	fmt.Println(people)
	f2(&people)
	fmt.Println(people)
	// 结构体指针1：使用new方法
	var peolpe1 = new(People)
	fmt.Printf("%T\n", peolpe1)
	// 结构体指针2：初始化
	var people2 = People{
		name:   "haha",
		gender: "男",
		age:    18,
	}
	fmt.Printf("%#v\n", people2)
	// 结构体中内存是连续的
	// 内存对齐？
}
