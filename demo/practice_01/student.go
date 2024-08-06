package main

import "fmt"

// 结构体 学生
type student struct {
	id    int // 学号是唯一的
	name  string
	class string
}

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学生管理-> 操作类
type managerStudent struct {
	allstudents []*student
}

func newstudentmanager() *managerStudent {
	return &managerStudent{
		allstudents: make([]*student, 0, 100), // 初始化100个
	}
}

// 添加学生
func (s *managerStudent) addStudent(newstu *student) {
	s.allstudents = append(s.allstudents, newstu)
}

// 修改学生信息
func (s *managerStudent) upgradeStudent(newstu *student) {
	for i, v := range s.allstudents {
		if newstu.id == v.id {
			s.allstudents[i] = newstu // 根据切片的索引直接把新学生赋值进去
			return
		}
	}
	fmt.Println("该学号未找到")
}

// 展示学生信息
func (s *managerStudent) showStudent() {
	for _, v := range s.allstudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s", v.id, v.name, v.class)
	}
}
