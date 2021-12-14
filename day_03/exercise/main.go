package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	// 遍历字符串切片，拿到每个名字
	for _, s := range users {
		distribution[s] = 0
		for _, c := range s {
			// 遍历每个字符串，拿到每个字符做判断
			if c == 'e' || c == 'E' {
				distribution[s] += 1
				coins -= 1
			} else if c == 'i' || c == 'I' {
				distribution[s] += 2
				coins -= 2
			} else if c == 'o' || c == 'O' {
				distribution[s] += 3
				coins -= 3
			} else if c == 'u' || c == 'U' {
				distribution[s] += 4
				coins -= 4
			}
		}
		fmt.Printf("%s拿到%d枚金币\n", s, distribution[s])
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
