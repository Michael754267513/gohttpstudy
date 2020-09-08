package main

import "fmt"

//传输值给ch 然后关闭channel
func Send(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

// 从接受channel里面获取值，传送到channel 2 里面。
func Change(ch1 <-chan int, ch chan<- int) {
	for {
		i, ok := <-ch1
		if !ok {
			break
		}
		ch <- i * i
	}
	close(ch)
}

func Get(ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Send(ch1)
	go Change(ch1, ch2)
	Get(ch2)
}
