package main

import "fmt"

// 常量
// const pi = 3.1415926

// const (
// 	statusOK = 200
// 	notFound = 404
// )

// 常见面试题 👇👇👇
// 批量声明常量时，如果某一行没有声明值，则缺省和上一常量值相同
// const (
// 	n1 = 100
// 	n2
// 	n3
// )

// iota: 类似枚举
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

const (
	b1 = iota // 0
	b2        // 1
	_         // 2
	b3        // 3
)

const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4        // 3
)

// 多个常量声明在一行
// iota: 每新增 “一行” 常量声明，iota+1
const (
	d1, d2 = iota + 1, iota + 2 // d1:1, d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2, d4:3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 这里指 二进制中 1 左移的位数（即2的10次方）
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)

// 定义了常量后，不能修改
// 程序运行期间，不会改变的量
func main() {
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(d1, d2, d3, d4)
	fmt.Println(KB, MB, GB)
}
