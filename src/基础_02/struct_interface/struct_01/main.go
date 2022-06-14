package main

import "fmt"

// 匿名字段
// 不常用！！！且同类型只能写一个！
// type person struct {
// 	string
// 	int
// }

// 结构体嵌套
type class struct {
	cls_name string
	stu_num  int
}

type class_teacher struct {
	name string
	age  int
}

type student struct {
	class               // 匿名嵌套结构体(匿名嵌套后，main函数中调用就可以直接.cls_name这样)
	cls_t class_teacher // 都有name字段，冲突了就不能这样写了，必须写全！
	name  string
	age   int
}

func main() {
	stu_1 := student{
		name: "张三",
		age:  18,
		class: class{
			cls_name: "1班",
			stu_num:  40,
		},
		cls_t: class_teacher{
			name: "1班班主任",
			age:  25,
		},
	}
	fmt.Println(stu_1.name, stu_1.cls_name, stu_1.cls_t.name)
}
