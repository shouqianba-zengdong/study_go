package main

import (
	"fmt"
)

// 切片练习题

func main() {
	s1 := make([]int, 5, 10)
	fmt.Println(s1)
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
	}
	fmt.Println(s1)
}
