package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOK bool

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOK bool   // false
)

// Go语言中，变量声明必须使用，不使用就会编译报错

// Go语言没有什么缩进的概念
func main() {
	age = 16
	name = "张三"
	isOK = true
	fmt.Print(isOK) // 在终端中输出打印的内容
	fmt.Println()
	fmt.Printf("name: %s\n", name) // %s 占位符
	fmt.Println(age)               // 打印内容后换行

	// 声明变量的同时，进行赋值
	var str1 string = "李四"
	fmt.Println(str1)

	// 类型推导（根据值判断该变量是什么类型）
	var str2 = "20"
	fmt.Println(str2)

	// 简短变量声明（相当于str2的声明方法），只能在函数中使用
	str3 := "haha"
	fmt.Println(str3)

	// 同一作用域内，变量不能重复声明
	// str1 := "zhangsan"

	// 匿名变量是一个特殊的变量，相当于一个黑洞，用于接收不使用的返回变量
}
