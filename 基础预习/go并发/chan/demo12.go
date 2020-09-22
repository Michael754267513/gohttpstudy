package main

import (
	"fmt"
)

var c = make(chan int, 5)

func f() {
	c <- 1

}

func main() {
	c <- 2

	go f()

	for {
		i, ok := <-c

		if !ok {

			break
		}
		fmt.Println(i, ok)
	}
	close(c)

}
