package main

import (
	"fmt"
)

func main() {
	b1 := true
	var b2 bool
	fmt.Printf("%T %v\n", b1, b1)
	fmt.Printf("%T %v", b2, b2)
}
