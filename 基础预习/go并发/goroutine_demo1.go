package main

import (
	"fmt"
	"runtime"
	"sync"
)

var ws = sync.WaitGroup{}

func print(i int) {
	fmt.Println(i)
	ws.Done()
}
func main() {
	runtime.GOMAXPROCS(2) // 限制并行执行的数量
	for i := 0; i < 100; i++ {
		ws.Add(1)
		go print(i) // goroutine 并发随机执行
	}
	ws.Wait()
}
