package main

import "fmt"

func main() {
	age := 19
	// if age > 20 {
	// 	fmt.Println("YES")
	// } else {
	// 	fmt.Println("NO")
	// }

	// 多个判断条件
	if age = 16; age > 30 { // 该临时赋值只在if判断语句中生效
		fmt.Println("中年")
	} else if age < 18 {
		fmt.Println("未成年")
	} else if age > 18 && age < 30 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习")
	}

}
