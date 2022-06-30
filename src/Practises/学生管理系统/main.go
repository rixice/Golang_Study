package main

import "fmt"

type student struct {
	id   int
	name string
}

var (
	allStudent map[int]*student // 声明变量
)

func all_student() {
	for k, v := range allStudent {
		fmt.Printf("学号：%d\t姓名：%s\n", k, v.name)
	}
}

func add_student(id int, name string) {
	newStu := newStudent(id, name) // 调用构造函数
	// 追加到allStudent这个map中
	allStudent[id] = newStu
}

func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func main() {
	// 1. 打印菜单
	var choice int
	var stu_id int
	var stu_name string
	allStudent = make(map[int]*student, 64) //初始化，开辟内存空间
	for {
		fmt.Println("##########欢迎来到学生管理系统##########")
		fmt.Println("1.查看所有学生\n2.新增学生\n3.删除学生\n0.退出系统")
		fmt.Printf("请输入你的操作：")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			all_student()
		case 2:
			fmt.Printf("请输入学生的id：")
			fmt.Scanln(&stu_id)
			if len(allStudent) == 0 {
				fmt.Printf("请输入学生的姓名：")
				fmt.Scanln(&stu_name)
				add_student(stu_id, stu_name)
			} else {
				for k, _ := range allStudent {
					if stu_id == k {
						fmt.Println("已存在该ID的学生！！！")
					} else {
						fmt.Printf("请输入学生的姓名：")
						fmt.Scanln(&stu_name)
						add_student(stu_id, stu_name)
					}
				}
			}
		case 3:
			if len(allStudent) == 0 {
				fmt.Printf("当前没有学生可以删除！！！")
			} else {
				fmt.Printf("请输入学生的id：")
				fmt.Scanln(&stu_id)
				for _, v := range allStudent {
					if stu_id == v.id {
						delete(allStudent, stu_id)
					} else {
						fmt.Println("不存在该名学生，请检查！！！")
					}
				}
			}
		case 0:
			return
		default:
			fmt.Println("ERROR:请正确输入选项！！！")
		}
	}
}
