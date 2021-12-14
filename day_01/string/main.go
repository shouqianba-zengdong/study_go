package main

import (
	"fmt"
	"strings"
)

func main() {
	// \具有特殊意义，转义符
	// 字符串只能用双引号
	path := "\"/Users/macuser/go/src/wosai-inc.com/study_go/fmt/main.go\""
	fmt.Printf("%s\n", path)
	// 多行字符串
	s := `
	天苍苍，雨茫茫
	风吹草低见牛羊`
	fmt.Printf("%s\n", s)

	// 字符串相关操作
	fmt.Println(len(s))

	// 字符串拼接
	ss := path + s
	fmt.Printf("%s\n", ss)
	fmt.Printf("%s%s\n", path, s)
	ss1 := fmt.Sprintf("%s%s", path, s)
	fmt.Printf("%s\n", ss1)

	// 字符串分割
	ss2 := strings.Split(path, "/")
	fmt.Println(ss2)

	// 包含
	fmt.Println(strings.Contains(s, "天苍苍"))
	fmt.Println(strings.Contains(s, "雨濛濛"))

	// 前缀
	fmt.Println(strings.HasPrefix(s, "天苍苍"))
	// 后缀
	fmt.Println(strings.HasSuffix(s, "牛羊"))

	// 子串出现的位置
	fmt.Println(strings.Index(s, "雨"))

	// 字符串拼接
	ss3 := strings.Join(ss2, "|")
	fmt.Println(ss3)
}
