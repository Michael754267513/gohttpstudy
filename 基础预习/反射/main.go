package main

import (
	"fmt"
	"reflect"
)

type Popoler struct {
}

func main() {
	var type_a int = 10
	var type_b interface{}
	var p Popoler
	// 反射出类型 类型（Type）和种类（Kind）
	fmt.Println(reflect.TypeOf(type_a))
	fmt.Println(reflect.TypeOf(type_b))
	fmt.Println(reflect.TypeOf(p).Kind())
	// 返回值
	fmt.Println(reflect.ValueOf(type_a))

	//    IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())

}
