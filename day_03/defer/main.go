package main

import (
	"fmt"
)

// defer关键字

func deferDemo() {
	// defer关键字用于资源释放
	fmt.Println("start")
	// defer关键字把它后面的语句延迟到函数即将返回的时候再执行
	// 当有多个defer关键字时，后进先出（栈）
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("哈哈哈")
	defer fmt.Println("哦哦哦")
	fmt.Println("end")
}

func main() {
	deferDemo()
}
