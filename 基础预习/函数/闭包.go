package main

import (
	"fmt"
)

// 闭包 变量在闭包里面，每次实例化都会保存 下次调用值会会根绝上次调用的值错位初始值
func add() func(int) int {
	var i int = 0
	return func(a int) int {
		i++
		//fmt.Println("i: ", i)
		return a + i
	}
}

func main() {
	//缩小变量作用域，减少对全局变量的污染
	//var wg sync.WaitGroup
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		fmt.Println(i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	a := add()
	fmt.Println(a(1))
	fmt.Println(a(1))
	b := add()
	fmt.Println(b(1))
	/*
		2
		3
		2
	*/

}
