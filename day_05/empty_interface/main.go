package main

import "fmt"

func show_type(a interface{}) {
	fmt.Printf("%T\n", a)
}

func main() {
	show_type("哈哈哈")
	show_type(123)
	show_type(uint32(122))
	show_type([]int{1, 2, 3, 4})
	show_type(map[string]interface{}{"name": 123, "age": 18})
}
