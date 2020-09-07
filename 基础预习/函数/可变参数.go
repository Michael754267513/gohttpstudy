package main

import (
	"fmt"
	"reflect"
)

// 切片类型  可变参数
func test(i ...int) {
	fmt.Println(reflect.TypeOf(i))
}

func ReturnStr(s ...string) (rs []string) {
	return s
}

func GetStr(s string) string {
	return s
}

func Adv(s string, av func(string) string) string {
	return av(s)
}

func main() {

	test(1)
	fmt.Println(ReturnStr("1", "2"))
	fmt.Println(Adv("aaaaa", GetStr)) // 函数作为一个变量传输
}
