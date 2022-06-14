package main

import "fmt"

// 结构体模拟实现其他语言中的“继承”

type animal struct {
	name string
}

// 方法
func (a animal) move() {
	fmt.Printf("%s会动", a.name)
}

// 类
type dog struct {
	feet int
	animal
}

// 给dog实现"汪汪汪"的方法
func (d dog) wang() {
	fmt.Printf("%s在：汪汪汪~\n", d.name) //dog中没有name，去嵌套的animal中找
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "二哈",
		},
	}
	d1.wang()
	d1.move()
}
