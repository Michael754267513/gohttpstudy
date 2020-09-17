package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(map[string]int)
	go func() { //开一个goroutine写map
		for j := 0; j < 10000; j++ {
			c[fmt.Sprintf("%d", j)] = j
		}
	}()

	go func() { //开一个goroutine写map
		for j := 10000; j < 20000; j++ {
			c[fmt.Sprintf("%d", j)] = j
		}
	}()

	go func() { //开一个goroutine读map
		for j := 0; j < 20000; j++ {
			fmt.Println(c[fmt.Sprintf("%d", j)])
		}
	}()

	time.Sleep(time.Second * 2)
}
