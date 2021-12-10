package main

import "fmt"

// 方法

type Dog struct {
	name string
	age  int
}

// 结构体构造函数
func newDog(name string, age int) *Dog {
	return &Dog{
		name: name,
		age:  age,
	}
}

// 定义一个方法，指定方法的接收者
// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体变量类型，多用类型名首字母小写
func (d Dog) wang() {
	fmt.Printf("%s汪汪汪~\n", d.name)
}

// 指针接收者
func (d *Dog) guonian() {
	fmt.Println("过年啦")
	d.age += 1
}

// 什么时候使用指针接收者？
// 需要修改接收者的值
// 接收者是拷贝代价比较大的对象
// 保证一致性

func main() {
	d1 := newDog("小东", 18)
	d1.wang()
	fmt.Printf("看看现在我多少岁：%d\n", d1.age)
	d1.guonian()
	fmt.Println(d1.age)
}
