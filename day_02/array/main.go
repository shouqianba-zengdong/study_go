package main

import (
	"fmt"
)

// 数组

func main() {
	// 数组的类型和容量是标识数组的一部分
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%v\n", a1)
	fmt.Printf("a2:%v\n", a2)
	// 数组初始化
	// 如果不初始化，默认元素都是零值
	a3 := [3]bool{true, true, true}
	fmt.Println(a3)
	a4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a4)
	// 部分初始化
	a5 := [...]int{0: 1, 2: 1, 6: 1}
	fmt.Println(a5)
	// 遍历数组方式一：索引
	as1 := [...]string{"北京", "上海", "广州", "深圳"}
	for i := 0; i < len(as1); i++ {
		fmt.Printf("%s\t", as1[i])
	}
	fmt.Println()
	// 遍历数组方式二：range函数
	for _, j := range as1 {
		fmt.Printf("%s\t", j)
	}
	// 多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)
	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
