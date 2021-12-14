package main

import (
	"fmt"
)

// go语言中的return不是原子操作，分为两步：
// 一：给返回值赋值。
// 二：RET指令
// 如果函数中存在defer，在这两步之间

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值为匿名变量=x=5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值为x=5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值为y=x=5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
