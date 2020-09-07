package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {

		m[stu.name] = &stu // 指针地址
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	/*
		小王子 => 大王八
		娜扎 => 大王八
		大王八 => 大王八
	*/
}
