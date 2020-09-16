package main

import "fmt"

type P struct {
	Name string
	Age  int
}

func main() {

	o := []P{
		P{"chain1", 20},
		P{"chain2", 21},
		P{"chain3", 22},
	}

	for _, v := range o {
		v.Age = 18 //  v在栈，o在堆，每次赋值是把堆里面的o，拉倒栈里面
		//o[k].Age = 18 // 正确修改方式，直接修改堆里面的值
	}

	fmt.Println(o)
	/*
		[{chain1 20} {chain2 21} {chain3 22}]
	*/
}
