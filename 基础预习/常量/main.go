package main

import "fmt"

// 常量 设置不变的
type people string

const (
	oldman people = "老男人"
	child  people = "小孩"
)

func main() {
	fmt.Println(oldman)
	fmt.Println(child)
}
