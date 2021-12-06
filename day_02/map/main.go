package main

import (
	"fmt"
)

// map

func main() {
	// map是引用类型，赋值前必须初始化
	// 要估算好map的容量
	m1 := make(map[string]string, 10)
	m1["姓名"] = "曾东"
	m1["年龄"] = "21"
	m1["性别"] = "男"
	fmt.Println(m1)
}
