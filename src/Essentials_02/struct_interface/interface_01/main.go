package main

import "fmt"

// 引出接口的实例
// 接口是一种类型，是一种特殊的类型，它规定了变量有哪些方法
// 在编程中，会遇到以下场景：
// 我不关心一个变量时什么类型，我只关心能调用它的什么方法

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型
}

type cat struct {
}

type dog struct {
}

type person struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("啊啊啊")
}

func da(x speaker) {
	// 接受一个参数，传进来什么，就打印什么
	x.speak() //挨打就要叫
}

//
func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker // 定义一个接口类型：speaker 的变量 ss
	ss = c1
	ss = d1
	ss = p1
	ss.speak()
	//fmt.Println(ss)
}
