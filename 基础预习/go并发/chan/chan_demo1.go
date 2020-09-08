package main

import "fmt"

// 通道关闭可以继续接受值，当通道没有值的时候，会接受通道类型默认值

func main() {
	ch := make(chan int, 2) //带缓冲通道
	ch <- 20
	ch <- 20
	close(ch)
	x := <-ch
	fmt.Println(x)

	ch1 := make(chan int, 200)

	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)

	for {
		i, ok := <-ch1
		if !ok {
			break
		}
		fmt.Println(i)
	}

}
