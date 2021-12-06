package main

import (
	"fmt"
)

func main() {
	// 切片声明
	var a []int
	var b []string
	fmt.Println(a, b)
	// 只定义不初始化，未分配内存空间
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	a = []int{1, 2, 3}
	b = []string{"北京", "上海", "广州", "深圳"}
	fmt.Println(a, b)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	// 长度和容量
	fmt.Printf("len(a):%d, cap(a):%d\n", len(a), cap(a))
	fmt.Printf("len(b):%d, cap(b):%d\n", len(b), cap(b))
	// 切片声明并且初始化
	var d = []bool{false, true}
	fmt.Println(d)
	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s1 := a1[0:4]
	s2 := a1[:4]
	// 切片的容量是从切片的第一个元素到底层数组最后一个元素
	s3 := a1[4:]
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d, cap(s2):%d\n", len(s2), cap(s2))
	fmt.Printf("len(s3):%d, cap(s3):%d\n", len(s3), cap(s3))
	// make函数创建切片
}
