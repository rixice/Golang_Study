package main

import "fmt"

func main() {
	a := 100
	b := 20

	c := a + b
	fmt.Printf("a+b结果: %v\n", c)

	c = a * b
	fmt.Printf("a*b结果: %v\n", c)

	c = a / b
	fmt.Printf("a/b结果: %v\n", c)

	c = a - b
	fmt.Printf("a-b结果: %v\n", c)

	a--
	c = a
	fmt.Printf("a--结果: %v\n", c) // --和++不能用在表达式里面

	// 二进制运算

	a = 4 // 0100
	b = 8 // 1000

	fmt.Printf("(a | b): %b\n", (a | b))
	fmt.Printf("(a & b): %b\n", (a & b))
	fmt.Printf("(a ^ b): %b\n", (a | b))
	fmt.Printf("(a << 2): %b\n", (a << 2))
}
