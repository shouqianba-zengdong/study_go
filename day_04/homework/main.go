package main

import (
	"fmt"
	"os"
)

// 用面向对象的思维方式编写一个学生信息管理系统

// 定义一个map记录学生信息
var (
	allStudent map[int]*Student
	i          int
)

// 学生的信息有id、姓名、年龄、分数
type Student struct {
	id, age, score int
	name           string
}

// 定义一个学生结构体构造函数
func newStudent(id, age, score int, name string) *Student {
	return &Student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}

// 展示所有学生信息
func showStudentList() {
	if len(allStudent) == 0 {
		fmt.Println("目前没有学生")
	} else {
		for _, v := range allStudent {
			fmt.Printf("学生id：%d\t", v.id)
			fmt.Printf("学生姓名：%s\t", v.name)
			fmt.Printf("学生年龄：%d\t", v.age)
			fmt.Printf("学生成绩：%d\t\n", v.score)
		}
	}
}

func addStudent() {
	var (
		id    int
		name  string
		age   int
		score int
	)
	fmt.Print("请输入学生id：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入学生年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入学生成绩：")
	fmt.Scanln(&score)
	s := newStudent(id, age, score, name)
	allStudent[i] = s
	i += 1
}

func deleteStudent() {
	var yourname string
	fmt.Print("请输入要删除学生的姓名")
	fmt.Scanln(&yourname)
	for k, v := range allStudent {
		if v.name == yourname {
			delete(allStudent, k)
		}
	}
}

func main() {
	i = 1
	// 初始化map（开辟内存空间）
	allStudent = make(map[int]*Student, 50)
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
		fmt.Printf("你选择了%s\n", choiceMap[choice])
		// 3.执行对应函数
		switch choice {
		case 1:
			showStudentList()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("请输入正确的正整数")
		}
		fmt.Println("*********************************")
	}
}
