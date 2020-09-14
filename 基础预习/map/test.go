package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1new := new([]int)
	s1make := make([]int, 2)
	fmt.Printf("new创建的类型：%v, 切片的类型: %v", reflect.TypeOf(s1new).Kind(), reflect.TypeOf(s1new))
	fmt.Printf("\nmake创建的类型：%v, 切片的类型: %v", reflect.TypeOf(s1make).Kind(), reflect.TypeOf(s1make))
	//new创建的类型：ptr, 切片的类型: *[]int
	//make创建的类型：slice, 切片的类型: []int

}
