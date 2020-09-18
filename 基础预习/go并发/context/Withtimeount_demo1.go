package main

import (
	"context"
	"fmt"
	"time"
)

//WithTimeount 的示例, 传递具有超时的 Context 以告知阻塞函数，它将在超时过后丢弃其工作。

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	start := time.Now()
	defer cancel()
	time.Sleep(time.Second * 6)
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			end := time.Now()
			fmt.Println(end.Sub(start))
			return
		}
	}

}
