package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileObj := path.Base(file)
	fmt.Println(funcName)
	fmt.Println(fileObj)
	fmt.Println(line)
}
