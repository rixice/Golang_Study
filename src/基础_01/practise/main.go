package main

import "fmt"

// 九九乘法表
func main() {
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%d x %d = %d\t", m, n, m*n)
		}
		fmt.Println()
	}
}
