package main

import "fmt"

// 多个类型可以实现同一个接口
// 一个类型可以实现多个接口

// 使用值接收者和指针接收者的区别？
type animal interface {
	eat(string)
	move()
}

type cat struct {
	name string
}

// 使用指针接收者：只能存结构体指针类型
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s", food)
}

func (c *cat) move() {
	fmt.Println("猫会跑~")
}

////////////////////////////////////
type chicken struct {
	name string
}

// 使用值接收者：结构体类型和结构体指针类型都能存
func (k chicken) eat(food string) {
	fmt.Printf("鸡吃%s", food)
}

func (k chicken) move() {
	fmt.Println("鸡会跑~")
}

func main() {
	var a animal

	c1 := cat{
		name: "蓝白英短",
	} // cat

	c2 := &cat{
		name: "暹罗猫",
	} // *cat

	a = &c1 // 实现animal这个接口的是cat的指针类型
	fmt.Println(a)
	a = c2
	fmt.Println(a)
}
