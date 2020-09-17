package main

import "fmt"

func main() {
	ch := make(chan int)
	defer close(ch) // 函数结束关闭chan
	go func() {
		ch <- 100 // chan赋值
	}()
	fmt.Println(<-ch) //获取chan的值
}
