package main

import "fmt"

// fmt占位符
func main() {
	var n = 100
	// 查看类型
	fmt.Printf("%T\n", n)

	// 查看变量的值
	fmt.Printf("%v\n", n)

	// 查看各种进制
	fmt.Printf("%d\n", n) //十进制
	fmt.Printf("%b\n", n) //二进制
	fmt.Printf("%o\n", n) //八进制
	fmt.Printf("%x\n", n) //十六进制

}
