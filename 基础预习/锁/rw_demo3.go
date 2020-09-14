package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var count int
var rw sync.RWMutex

func main() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}
func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
	/*
		goroutine 4 进入写操作...
		goroutine 4 写入结束，新值为：81
		goroutine 0 进入读操作...
		goroutine 0 读取结束，值为：81
		goroutine 1 进入读操作...
		goroutine 1 读取结束，值为：81
		goroutine 4 进入读操作...
		goroutine 4 读取结束，值为：81
		goroutine 3 进入读操作...
		goroutine 3 读取结束，值为：81
		goroutine 2 进入读操作...
		goroutine 2 读取结束，值为：81
		goroutine 0 进入写操作...
		goroutine 0 写入结束，新值为：887
		goroutine 1 进入写操作...
		goroutine 1 写入结束，新值为：847
		goroutine 2 进入写操作...
		goroutine 2 写入结束，新值为：59
		goroutine 3 进入写操作...
		goroutine 3 写入结束，新值为：81

	*/
}
