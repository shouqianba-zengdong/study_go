package main

import (
	"fmt"
)

const (
	pi  = 3.1415926
	day = 7
)

// 	批量声明常量时，如果某一行声明后没有赋值，默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota在const关键字出现的时候被重置为零
// const中每增加一行常量声明，iota的值加一
const (
	a1 = iota
	a2
	a3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(pi, day)
	fmt.Println(n1, n2, n3)
	fmt.Println(a1, a2, a3)
}
