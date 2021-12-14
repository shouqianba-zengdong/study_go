package main

import (
	"fmt"
	"os"
)

// 学生管理系统，结构体版
// 使用“面向对象”的思维方式编写一个学生信息管理系统。
// 学生有id、姓名、年龄、分数等信息
// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能

// 定义学生结构体，存储学生信息
type Student struct {
	id, age, score int
	name           string
}

// 定义班级结构体，结构体嵌套，保存所有学生信息
type Class struct {
	name    string
	student []*Student
}

// 学生结构体构造函数
func newStudent(id, age, score int, name string) *Student {
	return &Student{
		id:    id,
		age:   age,
		score: score,
		name:  name,
	}
}

// 班级结构体构造函数
func newClass(name string, student []*Student) *Class {
	return &Class{
		name:    name,
		student: student,
	}
}

func showStudentList(class *Class) {
	if len(class.student) == 0 {
		fmt.Println("尚无学生，请添加学生信息")
		return
	}
	for _, v := range class.student {
		fmt.Printf("学号：%d 姓名：%s 年龄：%d 成绩：%d\n", v.id, v.name, v.age, v.score)
	}
}

func addStudent(class *Class) {
	var (
		id, age, score int
		name           string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入成绩")
	fmt.Scanln(&score)
	s := newStudent(id, score, age, name)
	class.student = append(class.student, s)
}

func deleteStudent(class *Class) {
	var (
		deleteId int
		flag     bool
	)
	fmt.Print("请输入要删除学生的学号：")
	fmt.Scanln(&deleteId)
	for k, v := range class.student {
		if deleteId == v.id {
			if k == 0 {
				class.student = class.student[1:]
			} else {
				class.student = append(class.student[:k], class.student[k:]...)
			}
			flag = true
		}
	}
	if flag {
		fmt.Printf("学号为%d的学生信息已被删除\n", deleteId)
	} else {
		fmt.Printf("学号为%d的学生不存在\n", deleteId)
	}
}

func main() {
	myclass := Class{
		name:    "我的班级",
		student: make([]*Student, 0, 50),
	}
	var choiceMap = make(map[int]string, 10)
	choiceMap[1] = "查看所有学生"
	choiceMap[2] = "新增学生"
	choiceMap[3] = "删除学生"
	choiceMap[4] = "退出系统"
	// 1.打印菜单：展示学生管理系统提供给用户的功能
	for {
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(`
			1.查看所有学生
			2.新增学生
			3.删除学生
			4.退出
			`)
		// 2.等待用户选择要做什么
		fmt.Print("请输入你要进行的操作：")
		var choice int
		fmt.Scanln(&choice)
		_, ok := choiceMap[choice]
		if ok {
			fmt.Printf("你选择了%s\n", choiceMap[choice])
		} else {
			fmt.Println("请输入正确的正整数")
		}
		// 3.执行对应函数
		switch choice {
		case 1:
			showStudentList(&myclass)
		case 2:
			addStudent(&myclass)
		case 3:
			deleteStudent(&myclass)
		case 4:
			os.Exit(1)
		default:
			fmt.Println("请输入正确的正整数")
		}
		fmt.Println("*********************************************")
	}
}
