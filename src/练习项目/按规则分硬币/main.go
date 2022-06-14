package main

import (
	"fmt"
)

/*
5000枚金币，分配给users
名字中含'e'或'E'—— 1枚
名字中含'i'或'I'—— 2枚
名字中含'o'或'O'—— 3枚
名字中含'u'或'U'—— 4枚
计算每个用户分到多少金币，最后剩余？
*/

var (
	coins = 5000
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana",
		"Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() int {
	for _, name := range users {
		for _, check := range name {
			switch check {
			case 'e', 'E':
				distribution[name] += 1
				coins -= 1
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			default:
				distribution[name] += 0
			}
		}
		fmt.Printf("%v的金币：%d\n", name, distribution[name])
	}
	return coins
}
