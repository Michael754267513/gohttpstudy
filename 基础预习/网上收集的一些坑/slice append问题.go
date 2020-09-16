package main

import "fmt"

func main() {

	s := make([]string, 0, 3)
	s = append(s, "a", "b", "c")
	s2 := make([][]string, 0, 3)
	s2 = append(s2, s)
	s2[0][1] = "sss" // 修改s2的值，源切片的值也会变化
	fmt.Println(s2)
	fmt.Println(s)
	/*
		[[a sss c]]
		[a sss c]
	*/
}
