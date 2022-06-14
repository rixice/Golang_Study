package main

import "fmt"

// 函数：一段代码的封装
// 在一个命名的函数中不能够再声明命名函数
// defer语句会将其后面跟随的语句进行延迟处理

// deferDemo()
// a := f1
// b := f2
// c := f3
// fmt.Printf("%T\n", a)
// fmt.Printf("%T\n", b)
// fmt.Printf("%T\n", c)
// f7 := f5(f2)
// fmt.Printf("%T\n", f7)

// func deferDemo() {
// 	fmt.Println("hello")
// 	defer fmt.Println("world") // 延迟到函数即将返回的时候再执行
// 	defer fmt.Println("!!!")   // 类似于栈，先defer的放最后
// 	fmt.Println("？？？")
// }
////////////////////////////////////////////////
// Go语言中函数的return不是原子操作，在底层是分为两步来执行的
// 第一步：返回值赋值
// defer
// 第二步：真正的RET返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间！
////////////////////////////////////////////////
// func f1() int {
// 	x := 5
// 	defer func() {
// 		x++
// 	}()
// 	return x // x=5
// }

// func f2() (x int) {
// 	defer func() {
// 		x++
// 	}()
// 	return 5 // x=6
// }
//////////////////////////////////////////////////
// 变量的作用域
// 函数中查找变量的顺序
// 1. 先在函数内部查找
// 2. 找不到就往函数的外面查找（不一定要先命名，如下）
// 3. 一直找到全局为止
// func f1() {
// 	fmt.Println(x)
// }

// var x = 100

// func f2(x int) {
// 	fmt.Println(x)
// }
////////////////////////////////////////////////////
// 函数类型
// func f1() {
// 	fmt.Println("Hello world")
// }

// func f2() int {
// 	return 5
// }

// // 函数也可以作为参数的类型
// func f3(x func() int) {
// 	ret := x()
// 	fmt.Println(ret)
// }

// func ff(a, b int) int {
// 	return a + b
// }

// // 函数还可以作为返回值
// func f5(x func() int) func(int, int) int {
// 	return ff
// }
///////////////////////////////////////////
// f1(10, 20)
// 如果只是调用一次的函数，还可以简写成立即执行函数
// func(x, y int) {
// 	fmt.Println("Hello world")
// 	fmt.Println(x + y)
// }(100, 200)
///////////////////////////////////////////
// 闭包
// 闭包是一个函数，这个函数包含了他外部作用域的一个变量
// 底层原理：
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序，先在自己内部找，找不到往外层找
func main() {
	ret := f3(20, 30, f2)
	f1(ret)
}

func f1(f func()) {
	fmt.Println("This is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("This is f2")
	fmt.Println(x + y)
}

func f3(x, y int, f func(int, int)) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

// // 定义一个函数，对f2进行包装
// func f3(f func(int, int)) func() {
// 	tmp := func(){
// 		fmt.Println("hello")
// 	}
// 	return tmp
// }

// 匿名函数
// var f1 = func(x, y int) {
// 	fmt.Println(x + y)
// }
