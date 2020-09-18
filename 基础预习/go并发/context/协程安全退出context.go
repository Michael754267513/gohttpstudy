package main

import (
	"context"
	"time"
)

/*
主要用于在 goroutine 之间传递取消信号、超时时间、截止时间以及一些共享的值等
 context是协程并发安全的
 context包可以从已有的context实例派生新的context,形成context的树状结构,只要一个context取消了,派生出来的context将都会被取消

使用规则:

    不要将 Context放入结构体，Context应该作为第一个参数传入，命名为ctx。
	即使函数允许，也不要传入nil的 Context。如果不知道用哪种 Context，可以使用context.TODO()
	使用context的Value相关方法,只应该用于在程序和接口中传递和请求相关数据，不能用它来传递一些可选的参数
	context是协程安全的,可以传递不同的goroutine
*/
func work(ctx context.Context, msg string) {
	for {
		select {
		case <-ctx.Done():
			println(msg, "goroutinue is finish......")
			return
		default:
			println("goroutinue is running", time.Now().String())
			time.Sleep(time.Second)
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "withCancel")
	time.Sleep(time.Second * 3)
	println("cancel......")
	cancel()
	time.Sleep(time.Second * 3)
	println("finish")
}
