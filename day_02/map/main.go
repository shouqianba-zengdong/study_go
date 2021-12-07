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
	// 如果没有key返回零值，但不推荐
	// 约定俗成：如果返回值是bool类型，则用ok接受返回值
	v, ok := m1["爱好"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此key")
	}
	// map的遍历
	for k, v := range m1 {
		fmt.Printf("key:%s value:%s\n", k, v)
	}
	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
	// 删除
	delete(m1, "性别")
	fmt.Println(m1)
}
