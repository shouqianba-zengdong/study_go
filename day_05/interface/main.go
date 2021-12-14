package main

import "fmt"

// 接口：一个变量实现了接口中定义的所有方法，就可以说这个变量实现了该接口

type cat struct {
	name string
}

type dog struct {
	name string
}

func (c cat) speak() {
	fmt.Printf("%s：喵喵喵\n", c.name)
}

func (d dog) speak() {
	fmt.Printf("%s：汪汪汪\n", d.name)
}

func da(s speaker) {
	s.speak()
}

type speaker interface {
	speak()
}

func main() {
	c1 := cat{
		name: "小喵",
	}
	d1 := dog{
		name: "小汪",
	}
	da(c1)
	da(d1)
}
