package main

import (
	"fmt"
)

// append方法为切片追加元素

func main() {
	s1 := []string{"北京", "上海", "广州"}
	fmt.Printf("%v len:%d cap:%d\n", s1, len(s1), cap(s1))
	// 调用append方法必须使用原油变量接受返回值
	s1 = append(s1, "深圳")
	// append追加元素，原来数组放不下，会把底层数组换一个
	fmt.Printf("%v len:%d cap:%d\n", s1, len(s1), cap(s1))
	// 切片容量升级策略:新容量大于原来两倍，新容量
	// 原容量小于1024,最终容量翻倍
	// 原容量大雨1024,1/4扩容
	fmt.Println(s1)
	s2 := []string{"武汉", "长沙"}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
