package main

import (
	"fmt"
)

// make函数创建切片

func main() {
	// 切片的本质是一块连续的内存
	s1 := make([]int, 5, 10)
	fmt.Printf("s1:%v len:%d cap:%d\n", s1, len(s1), cap(s1))
	// 切片的赋值
	s2 := []int{1, 3, 5}
	s3 := s2
	fmt.Println(s2, s3)
	s3[0] = 100
	fmt.Println(s2, s3)
	// 切片的遍历
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%d\t", s2[i])
	}
	for _, v := range s3 {
		fmt.Printf("%d\t", v)
	}
}
