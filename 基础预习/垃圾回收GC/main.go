//GODEBUG=gctrace=1  可以查看GC 过程

package main

import (
	"fmt"
	_ "net/http/pprof"
	"time"
)

func main() {
	a := "123"
	fmt.Println(a)

	B := make(chan int, 8)
	go func() {
		B <- 1
	}()
	go func() {
		fmt.Println(<-B)
	}()

	c := make([]int, 10, 20)
	c = append(c, 1, 2, 3)
	fmt.Println(c)

	time.Sleep(1 * time.Second)
}
