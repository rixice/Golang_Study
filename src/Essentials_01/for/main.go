package main

import "fmt"

func main() {
	// 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	n := 1
	for n < 10 {
		fmt.Printf("%d", n)
		n++
	}

	fmt.Println("###################")

	// 变种1
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种2
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 无限循环
	// for{
	// 	fmt.Println("ok")
	// }

	// for range循环
	// s := "Hello World"
	// for i, v := range s {
	// 	fmt.Printf("%d,%c\n", i, v)
	// }

}
