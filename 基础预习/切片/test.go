package main

import "fmt"

func main() {
	//var a = make([]int, 5, 10)
	//for i := 0; i < 10; i++ {
	//	a = append(a, i)
	//}
	//a = append(a,2)
	//fmt.Println(a)
	//sort.Ints(a)
	//fmt.Println(a)

	//s := make([]int,0,1)
	s := []int{}
	if s == nil {
		fmt.Println(" nil 切片为空")
	}
	if len(s) == 0 {
		fmt.Println(" len 切片为空")
	}
}
