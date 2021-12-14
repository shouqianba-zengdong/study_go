package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 计算字符串中汉字的数量

func main() {
	s1 := "hello曾东你好haha"
	sum := 0
	// 遍历字符串
	for _, v := range s1 {
		// 判断字符是否为汉字
		if unicode.Is(unicode.Han, v) {
			// 是汉字则总数加一
			sum += 1
		}
	}
	fmt.Printf("\"%s\"中有%d个汉字\n", s1, sum)
	// 计算单词出现的次数
	s2 := "how do you do do"
	// 切割字符串得到切片
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, len(s3))
	// 遍历切片，把单词存入map
	for _, w := range s3 {
		// 判断map中是否存在该key，不存在key的值置为1，存在key的值加一
		_, ok := m1[w]
		if !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	for k, v := range m1 {
		fmt.Printf("%s出现的次数为%d\n", k, v)
	}
	// 回文数判断
	s4 := "123321"
	for i := 0; i < len(s4)/2; i++ {
		if s4[i] != s4[len(s4)-i-1] {
			fmt.Println("非回文数")
			return
		}
	}
	fmt.Println("回文数")
	// 回文数判断，双指针
	s5 := "12345432"
	for i, j := 0, len(s5)-1; i < j; i, j = i+1, j-1 {
		if s5[i] != s5[j] {
			fmt.Println("非回文数")
			return
		}
	}
	fmt.Println("回文数")
}
