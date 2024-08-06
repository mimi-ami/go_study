package main

import (
	"fmt"
	"go_study/demo/calc"
)

// 先执行go build  会生成一个.exe文件，然后执行这个.exe文件
// 学员信息管理系统
// 练习题

// func showMenu() {
// 	fmt.Println("欢迎来到学生信息管理系统")
// 	fmt.Println("1. 添加学生")
// 	fmt.Println("2. 编辑学生")
// 	fmt.Println("3. 展示所有学生信息")
// 	fmt.Println("4. 退出")
// }

// // 获取用户输入的学生信息
// func getInput() *student {
// 	var (
// 		id    int
// 		name  string
// 		class string
// 	)
// 	fmt.Print("请按照要求输入学生信息：")
// 	fmt.Print("请输入学号：")
// 	fmt.Scanln(&id)
// 	fmt.Print("请输入姓名：")
// 	fmt.Scanln(&name)
// 	fmt.Print("请输入班级：")
// 	fmt.Scanln(&class)
// 	stu := newStudent(id, name, class)
// 	return stu
// }

func main() {
	re := calc.Add(1, 2)
	fmt.Println(re)
	// sm := newstudentmanager()
	// for {
	// 	// 打印系统菜单
	// 	showMenu()
	// 	// 等待用户选择功能
	// 	var input int
	// 	fmt.Println("请选择：")
	// 	fmt.Scan(&input)
	// 	fmt.Println("用户选择的功能是：", input)
	// 	// 执行用户选择的功能
	// 	switch input {
	// 	case 1:
	// 		// 添加学生
	// 		stu := getInput()
	// 		sm.addStudent(stu)
	// 	case 2:
	// 		// 编辑学生信息
	// 		stu := getInput()
	// 		sm.upgradeStudent(stu)
	// 	case 3:
	// 		// 展示学生信息
	// 		sm.showStudent()
	// 	case 4:
	// 		// 退出
	// 		os.Exit(0)
	// 	}

	// }

}

	

