package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	log.SetOutput(fileObj)
	i := 100
	for {
		log.Println("这是一条测试日志")
		time.Sleep(3 * time.Second)
		if i > 100 {
			return
		}
	}
}
