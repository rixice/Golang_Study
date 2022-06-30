package main

import "fmt"

// 接口示例2

type animal interface {
	eat(string)
	move()
}

type cat struct {
	name string
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s", food)
}

func (c cat) move() {
	fmt.Println("猫会跑~")
}

////////////////////////////////////
type chicken struct {
	name string
}

func (k chicken) eat(food string) {
	fmt.Printf("鸡吃%s", food)
}

func (k chicken) move() {
	fmt.Println("鸡会跑~")
}

func main() {
	var a animal
	var c1 = cat{
		name: "暹罗猫",
	}
	var k1 = chicken{
		name: "铁公鸡",
	}

	a = c1
	a.eat("鱼")
	a.move()

	a = k1
	a.eat("饲料")
	a.move()
}
