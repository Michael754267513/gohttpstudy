package main

import "fmt"

// 通道关闭可以继续接受值，当通道没有值的时候，会接受通道类型默认值

func main() {

	ch1 := make(chan int, 200) // 带缓冲的通道

	for i := 0; i < 100; i++ {
		ch1 <- i // 传输值到chan
	}
	close(ch1)

	for {
		i, ok := <-ch1 // 读取chan 里面内容，如果有则是ok，如果没有则是flase
		if !ok {
			break
		}
		fmt.Println(i) // chan里面获取到的值
	}

}
