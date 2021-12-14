package main

import (
	"fmt"
)

// map和slice组合

func main() {
	// 元素为map的slice
	s1 := make([]map[int]string, 10, 10)
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "曾东"
	fmt.Println(s1)
	// 值为slice类型的map
	m1 := make(map[string][]int, 10)
	m1["学号"] = []int{10, 20, 30}
	fmt.Println(m1)
}
