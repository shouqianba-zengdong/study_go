package main

import "fmt"

// 结构体嵌套

type address struct {
	provice string
	city    string
}

type person struct {
	name string
	age  int
	addr address
}

type company struct {
	name int
	addr address
}

func main() {
	p1 := person{
		name: "曾东",
		age:  18,
		addr: address{
			provice: "湖南",
			city:    "娄底",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.addr.provice)
}
