package main

import "fmt"

type person struct {
	name  string
	age   int
	sex   string
	hobby []string
}

func main() {
	// panic 和 recover
	// funcA()
	// funcB()
	// funcC()
	///////////////////////////////
	// 递归: 自己调用自己
	////////////////////////////////
	// 自定义类型 和 类型别名
	// type 后面跟的是类型
	// type myInt int     // 自定义类型
	// type yourInt = int // 类型别名

	// var n myInt
	// n = 100
	// var m yourInt
	// m = 1000
	// fmt.Printf("%T\n", n)
	// fmt.Printf("%T\n", m)
	//////////////////////////////////
	// 结构体(是 值类型)
	var p person
	p.name = "张三"
	p.sex = "男"

	var p3 = &person{ // 结构体指针
		name: "李四",
		age:  18,
	}

	fmt.Printf("%p\n", &p3.name) // 结构体占用一块连续的内存
	fmt.Printf("%p\n", &p3.age)
	// f(p) // go语言中函数传参永远是拷贝
	// f2(&p)
	// fmt.Println(p.sex)
	// // 匿名结构体
	// var s struct {
	// 	x string
	// 	y int
	// } // 多用于临时场景
	// s.x = "hello"
	// s.y = 100

	// var n person
	// // 通过字段赋值
	// n.name = "zhangsan"
	// n.age = 18
	// n.gender = "男"
	// n.hobby = []string{"篮球", "足球"}
	// fmt.Println(n)
	// // 访问变量n的字段
	// fmt.Println(n.name)
}
func f(x person) {
	x.sex = "女"
}

func f2(x *person) {
	x.sex = "女" // 根据内存地址进行修改，修改的就是原来的变量
}

// func funcA() {
// 	fmt.Println("A")
// }

// func funcB() {
// 	// 刚刚打开一个数据库连接
// 	defer func() {
// 		error := recover() // 尝试恢复现场
// 		fmt.Println(error)
// 		fmt.Println("释放数据库连接...")
// 	}()
// 	panic("出现了严重的错误！！！") // 程序崩溃退出！
// 	fmt.Println("B")
// }

// func funcC() {
// 	fmt.Println("C")
// }
