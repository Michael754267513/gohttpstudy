package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}

	s1 = append(s1[:2], s1[3:]...) //删除index为2 的元素
	fmt.Printf("%v", s1[:])
	/*

		[1 2 4 5 6]

	*/

}
