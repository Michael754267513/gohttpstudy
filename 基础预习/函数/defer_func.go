package main

import "fmt"

func df() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func main() {
	fmt.Println(df())
	// defer 在return 后执行
	/*
	 6
	*/
}
