package main

import "fmt"

// 递归：函数自己调用自己
// 注意：递归一定要有一个出口，否则是死循环

func haha(n uint) uint {
	if n <= 1 {
		return 1
	} else {
		return n * haha(n-1)
	}
}

func main() {
	r := haha(5)
	fmt.Println(r)
	r2 := taijie(5)
	fmt.Println(r2)
}

// 上台阶面试题，有n层楼梯，一次可以走一步也可以走两步，问有多少种走法
func taijie(n uint) uint {
	// 如果只有一层台阶，那么只有一种走法
	if n == 1 {
		return 1
	} else if n == 2 {
		// 如果有两层台阶，那么有两种走法
		return 2
	} else {
		return taijie(n-1) + taijie(n-2)
	}
}
