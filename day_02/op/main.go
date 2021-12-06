package main

import (
	"fmt"
)

// 运算符

func main() {
	var (
		a = 5
		b = 2
	)
	// 算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在等号的右边赋值
	b++

	// 关系运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	// 逻辑运算符
	// &&与
	// ||或
	// ！非

	// 位运算符
	fmt.Println(5 & 2)  // 按位与
	fmt.Println(5 | 2)  // 按位或
	fmt.Println(5 << 1) // 二进制左移一位
	fmt.Println(5 >> 2) // 二进制右移一位
}
