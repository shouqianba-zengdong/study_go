package main

import "fmt"

// 结构体模拟实现其他语言的继承

type animal struct {
	name string
}

// 给动物实现一个移动的方法
func (a animal) move() {
	fmt.Printf("我是动物%s，我正在移动\n", a.name)
}

// 定义结构体：狗
type dog struct {
	feet   int
	animal //animal的方法，dog也能使用
}

// 给狗结构体实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("我是小狗%s，我会汪汪汪\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "旺财",
		},
	}
	d1.wang()
	d1.move()
}
