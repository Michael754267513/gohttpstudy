package main

import (
	"fmt"
	"reflect"
)

func main() {

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}

	// 切片类型无法单独做判断，可以根据deep equal的判断值是否一致
	if reflect.DeepEqual(s1, s2) {
		fmt.Println("true")
	} else {
		fmt.Println("flase")
	}

	// 判断空切片
	s1check := make([]int, 0)
	if len(s1check) == 0 {
		fmt.Printf("判断空切片\n")
	}

	fmt.Println(s1[:])
	/*

	 	true
		判断空切片
		[1 2 3 4 5]


	*/

}
