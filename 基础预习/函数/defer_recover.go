package main

import "fmt"

func ErrorRecover() {
	// 申明错误 defer表示函数尾部执行
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("GG")
		}
	}()
	panic("err")
}

func main() {

	ErrorRecover()
	/*
		GG
	*/
}
