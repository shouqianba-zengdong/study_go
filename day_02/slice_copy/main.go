package main

import (
	"fmt"
)

// copy函数:ctrl + c ctrl + v

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)
	// 从切片中删除元素：go语言没有内置删除切片元素方法
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)
}
