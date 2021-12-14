package main

import (
	"fmt"
)

// if条件判断
func main() {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("澳门首家线上赌场开业了")
	// } else {
	// 	fmt.Println("该写暑假作业了")
	// }
	if age := 19; age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好少年")
	}
}
