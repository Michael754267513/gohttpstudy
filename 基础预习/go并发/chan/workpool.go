package main

import (
	"fmt"
	"time"
)

/*
	在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨。
*/

func run(i int, ch1 chan<- int, ch2 <-chan int) {
	for j := range ch2 {
		fmt.Printf("worker:%d start job:%d\n", i, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", i, j)
		ch1 <- j * 2
	}
}

func main() {
	ch1 := make(chan int, 100) // 异步
	defer close(ch1)
	ch2 := make(chan int, 100)
	defer close(ch2)
	for i := 1; i <= 3; i++ {
		go run(i, ch1, ch2)
	}
	for i := 0; i < 100; i++ {
		ch2 <- i
	}

	for a := 1; a <= 100; a++ {
		<-ch1
	}

}
