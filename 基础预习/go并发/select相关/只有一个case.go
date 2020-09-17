package main

import "fmt"

func main() {

	ch := make(chan int)
	defer close(ch)
	go func() {
		for i := 0; i <= 1; i++ {
			ch <- i
		}
	}()

	for {
		select {
		case i, ok := <-ch:
			fmt.Printf("\n接收到的值：%v,是否ok：%v", i, ok)
		}
	}
	/*
		接收到的值：0,是否ok：false
		接收到的值：0,是否ok：false
	*/
}
