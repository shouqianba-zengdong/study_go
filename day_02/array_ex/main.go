package main

import (
	"fmt"
)

func main() {
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 8
	for i := 0; i < len(a1); i++ {
		target := sum - a1[i]
		for j := i + 1; j < len(a1); j++ {
			if a1[j] == target {
				fmt.Printf("(%d, %d)", i, j)
			}
		}
	}
}
