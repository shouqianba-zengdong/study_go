package main

import "fmt"

// 接口类型断言

func type_assert(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("%v is a %v\n", a, v)
	case int:
		fmt.Printf("%v is a %v\n", a, v)
	case bool:
		fmt.Printf("%v is a %v\n", a, v)
	default:
		fmt.Println("unsupport type")
	}
}

func main() {
	type_assert(100)
	type_assert("hahaha")
	type_assert(true)
	type_assert([...]int{1, 2, 3, 4})
}
