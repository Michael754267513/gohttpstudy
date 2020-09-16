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
	oPointer := make([]*P, 0, 3)
	for _, v := range o {
		oPointer = append(oPointer, &v) // v的值一直在变，但是指针的地址没有变，所以最后得到的值就是v最后一次的值
		//a := v
		//oPointer = append(oPointer, &a)  // 处理方式
		fmt.Println(oPointer)
	}
	for _, v := range oPointer {
		fmt.Println(v)
	}
	/*
		[0xc0000044c0 0xc0000044c0 0xc0000044c0]
		&{chain3 22}
		&{chain3 22}
		&{chain3 22}
	*/
}
