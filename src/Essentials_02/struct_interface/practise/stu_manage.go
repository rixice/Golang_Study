package main

import "fmt"

// 学生管理功能实现
type student struct {
	id   int
	name string
}

// 学生管理者
type studentMgr struct {
	allStudent map[int]*student
}

// 查看学生
func (s studentMgr) showStudent() {
	fmt.Println("#########################################")
	for _, v := range s.allStudent {
		fmt.Printf("学号: %d\t姓名: %s\n", v.id, v.name)
	}
	fmt.Println("#########################################")
}

// 修改学生
func (s studentMgr) editStudent() {
	var stu_id int
	var stu_name string
	fmt.Printf("请输入要修改的学生学号：")
	fmt.Scanln(&stu_id)
	fmt.Printf("请输入新的学生姓名：")
	fmt.Scanln(&stu_name)
	newStu := student{
		id:   stu_id,
		name: stu_name,
	}
	s.allStudent[stu_id] = &newStu
}

// 新增学生
func (s studentMgr) addStudent() {
	var stu_id int
	var stu_name string
	fmt.Printf("请输入学号: ")
	fmt.Scanln(&stu_id)
	fmt.Printf("请输入姓名: ")
	fmt.Scanln(&stu_name)
	newStu := student{id: stu_id, name: stu_name}
	s.allStudent[stu_id] = &newStu
}

// 删除学生
func (s studentMgr) delStudent() {
	var stu_id int
	fmt.Printf("请输入要删除的学生学号：")
	fmt.Scanln(&stu_id)
	delete(s.allStudent, stu_id)
}
