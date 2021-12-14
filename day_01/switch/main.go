package main

import (
	"fmt"
)

type Finger struct {
	name   string
	length int
	index  int
}

func main() {
	var finger Finger
	finger.index = 1
	// switch语句的作用：减少大量的判断语句
	switch finger.index {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}
}
