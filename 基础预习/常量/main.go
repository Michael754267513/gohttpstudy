package main

import "fmt"

// 常量 设置不变的
type people string

// 全局变量
var age int

const (
	oldman people = "老男人"
	child  people = "小孩"
)

func main() {
	// 赋值
	age = 10

	// 局部变量
	var name string
	name = "测试"
	fmt.Println(name)
	fmt.Println(oldman)
	fmt.Println(child)
}
