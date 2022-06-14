package main

import "fmt"

func main() {
	a := 10
	b := 5

	c := a > b
	fmt.Printf("C的值为：%v\n", c)

	c = a == b
	fmt.Printf("C的值为：%v\n", c)

	c = a != b
	fmt.Printf("C的值为：%v\n", c)

	t := true
	f := false
	c = t && f
	fmt.Printf("c的值为：%v\n", c)

	c = t || f
	fmt.Printf("c的值为：%v\n", c)

	c = !f
	fmt.Printf("c的值为：%v\n", c)
}
