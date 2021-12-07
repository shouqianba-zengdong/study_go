package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	// 取出map中所有的key存入切片keys
	keys := make([]string, 0, 200)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	// 对切片排序
	sort.Strings(keys)
	// 按照排序后的key遍历value
	for _, key := range keys {
		fmt.Printf("key:%s,value:%d\n", key, scoreMap[key])
	}
}
