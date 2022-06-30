package main

import "fmt"

// type x struct {
// 	a int8 // 8bit → 1byte
// 	b int8
// 	c int8
// 	d string
// }

func main() {
	// m := x{
	// 	a: int8(10),
	// 	b: int8(20),
	// 	c: int8(30),
	// 	d: "hello",
	// }
	// fmt.Printf("%p\n", &(m.a))
	// fmt.Printf("%p\n", &(m.b))
	// fmt.Printf("%p\n", &(m.d))
	////////////////////////////////////////////////////
	// p1 := newPerson("张三", 18)
	// p2 := newPerson("李四", 20)
	// fmt.Println(p1.name)
	// fmt.Println(p2.name)
	////////////////////////////////////////////////////
	// d1 := newDog("旺财")
	// d1.wang()
	p1 := person{
		name: "zhangsan",
		age:  18,
	}
	p1.new_year()
	fmt.Println(p1.age)
}

// 构造函数
// 约定俗成：用new开头
// 当结构体比较大的时候，尽量使用结构体指针，减少程序的运行开销
// func newPerson(name string, age int) *person {
// 	return &person{
// 		name: name,
// 		age:  age,
// 	}
// }
func newDog(name string) *dog {
	return &dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，一般用类型名首字母小写表示

// 值接收者: 传拷贝
func (d dog) wang() {
	fmt.Printf("%s: 汪汪汪~", d.name)
}

// 指针接收者: 传内存地址（当 需要修改接收值的时候/字段较多 时使用）
func (p *person) new_year() {
	p.age++
}

type person struct {
	name string
	age  int
}

// 方法
type dog struct {
	name string
}

// Go中，如果标识符的首字母是大写的，就表示对外部包可见（public）

// Dog 这是一个狗的结构体（go语言中要求public前要有注释）
type Dog struct {
	name string
}
