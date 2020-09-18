package main

import (
	"context"
	"fmt"
	"time"
)

//WithDeadline 的示例,通过一个截止日期的 Context 来告知一个阻塞的函数，一旦它到了最终期限，就放弃它的工作
// 处理多久直接退出

func main() {
	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	start := time.Now()

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()
	time.Sleep(time.Second * 6)
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			end := time.Now()
			fmt.Print(end.Sub(start))
			return
		}
	}

}
