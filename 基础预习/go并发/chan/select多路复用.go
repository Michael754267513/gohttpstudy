package main

import "fmt"

/*
使用select语句能提高代码的可读性。
    可处理一个或多个channel的发送/接收操作。
    如果多个case同时满足，select会随机选择一个。
    对于没有case的select{}会一直等待，可用于阻塞main函数。
*/

func main() {
	// 当i=0 时，通道没有值无法满足 x:=<-ch 进去第二个case，
	// 当i=1时，ch的值为0，满足第一个，因为通道的cap为1 无法在进行往通道赋值
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			fmt.Println("通道赋值")
		case x := <-ch:
			fmt.Println(x) // 0 2 4 6 8
		}
	}
}
