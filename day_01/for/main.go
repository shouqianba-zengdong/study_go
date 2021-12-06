package main

import (
	"fmt"
)

func main() {
	// for循环
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// 跳出for循环
	for k := 1; k < 10; k++ {
		if k == 5 {
			break
		}
		fmt.Println(k)
	}

	// 不执行某次循环
	for x := 1; x <= 10; x++ {
		if x == 5 {
			continue
		}
		fmt.Println(x)
	}
}
