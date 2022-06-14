package main

import "fmt"

// 整型

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)

	// 八进制（0~7）
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制（0~f）
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// 十进制 转 二进制
	fmt.Printf("%b\n", i1)
	// 十进制 转 八进制
	fmt.Printf("%o\n", i1)
	// 十进制 转 十六进制
	fmt.Printf("%x\n", i1)

	// 查看变量类型
	fmt.Printf("%T\n", i3)

	// 声明 int8 类型的变量
	i4 := int8(9) // 明确指定int8类型，否则缺省为int类型
	fmt.Printf("%T", i4)
}
