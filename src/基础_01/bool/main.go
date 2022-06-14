package main

import "fmt"

// 布尔值
func main() {
	b1 := true
	var b2 bool
	fmt.Printf("%T value: %v \n", b1, b1)
	fmt.Printf("%T value: %v \n", b2, b2) // 默认是false
}
