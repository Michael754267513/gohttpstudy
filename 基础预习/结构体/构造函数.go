package main

import "fmt"

type People struct {
	Name string
	Age  int
	sex  bool
}

// 构造函数返回一个结构体
func newPeople(name string, age int, sex bool) People {
	return People{
		Name: name,
		Age:  age,
		sex:  sex,
	}
}
func main() {
	p1 := newPeople("laoshu", 20, true)
	fmt.Printf("p1: %v", p1)
	//p1: {laoshu 20 true}

}
