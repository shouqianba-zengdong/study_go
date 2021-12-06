package main

import (
	"fmt"
)

// byte和rune类型

// go语言为了处理非ASCII码类型字符定义了新的rune类型

func main() {
	s := "Hello曾东이륙하다."
	// len求的是byte字节的数量
	fmt.Println(len(s))
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
}
