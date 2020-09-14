package main

import "fmt"

func main() {
	// 匿名函数多用于实现回调函数和闭包
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(10, 20))

}
