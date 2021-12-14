package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat()
}

type cat struct {
	name string
	feet uint8
}

type chicken struct {
	name string
	feet uint8
}

func (c cat) move() {
	fmt.Printf("%s有%d只脚能移动\n", c.name, c.feet)
}

func (c cat) eat() {
	fmt.Printf("%s在吃东西\n", c.name)
}

func (ch chicken) move() {
	fmt.Printf("%s有%d只脚能移动\n", ch.name, ch.feet)
}

func (ch chicken) eat() {
	fmt.Printf("%s在吃东西\n", ch.name)
}

func main() {
	c1 := cat{
		name: "小喵",
		feet: 4,
	}
	ch1 := chicken{
		name: "小鸡",
		feet: 2,
	}
	c1.eat()
	c1.move()
	ch1.eat()
	ch1.move()
}
