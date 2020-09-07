package main

import "fmt"

// 接口使用场景，支付方式有多种，用同一个接口去实现
type Action interface {
	Say()
}

type Cat struct {
}
type Dog struct {
}

func (action Cat) Say() {
	fmt.Println("i am cat")
}

func (action Dog) Say() {
	fmt.Println("i am dog")
}

func main() {
	cat := Cat{}
	dog := Dog{}
	var A Action
	A = cat
	A.Say()
	A = dog
	A.Say()

}
