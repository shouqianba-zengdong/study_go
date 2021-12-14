package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与字符串

// 1.把go语言中的结构体变量转化成json格式字符串
// 2.json格式字符串转化成go语言中能识别的结构体变量

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "曾东",
		Age:  18,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err: %v", err)
		return
	} else {
		// 序列化之后返回的是rune类型的切片
		fmt.Println(string(b))
	}
	// 反序列化
	str := `{"name": "曾东", "age": 10}`
	p2 := new(person)
	// 传指针是为了在json.unmarshal内部修改p2的值
	json.Unmarshal([]byte(str), p2)
	fmt.Println(*p2)
}
