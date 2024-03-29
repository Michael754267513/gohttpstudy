package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"
var key1 string = "name1"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	valueCtx1 := context.WithValue(ctx, key, "【监控2】")
	go watch1(valueCtx)
	go watch2(valueCtx1)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			go getwatch1(ctx)
			time.Sleep(2 * time.Second)
		}
	}
}
func getwatch1(ctx context.Context) {
	fmt.Println(ctx.Value(key), "---监控退出，停止了...")
}

func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
