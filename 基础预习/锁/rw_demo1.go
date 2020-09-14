package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int
	rwLock sync.RWMutex
)

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				rwLock.Lock()
				count++
				rwLock.Unlock()
			}
			fmt.Println(count)
		}()
	}
	time.Sleep(time.Second)
	//fmt.Scanf("\n") //等待子线程全部结束
}
