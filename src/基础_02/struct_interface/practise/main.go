package main

import "fmt"

var admin studentMgr

// 学生管理系统
func Menu() int {
	var choice int
	fmt.Println("欢迎来到学生管理系统")
	fmt.Println("1.查看所有学生\n2.新增学生\n3.修改学生\n4.删除学生\n0.退出系统")
	fmt.Printf("请输入你的选择：")
	fmt.Scanln(&choice)
	return choice
}

func main() {
	admin = studentMgr{
		allStudent: make(map[int]*student, 100),
	}
	for {
		choice := Menu()
		switch choice {
		case 1:
			admin.showStudent()
		case 2:
			admin.addStudent()
		case 3:
			admin.editStudent()
		case 4:
			admin.delStudent()
		case 0:
			return
		default:
			fmt.Println("不存在该选项！！！")
		}
	}
}
