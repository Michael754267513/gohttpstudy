package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ch := make(chan int)
	start := time.Now()
	defer cancel()
	defer close(ch)
	for {
		select {
		case <-ch:
			fmt.Println("Work")
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			end := time.Now()
			fmt.Println(end.Sub(start))
			return
		}
	}

}
