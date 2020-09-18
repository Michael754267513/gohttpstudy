package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mutex := sync.Mutex{}
	// 创建唤醒锁
	cond := sync.NewCond(&mutex)

	for i := 0; i < 10; i++ {
		go func(num int, cond *sync.Cond) {
			cond.L.Lock()
			if num < 3 {
				//小于3的休眠，等待唤醒
				cond.Wait()
			}
			fmt.Println("num", num, "time", time.Now().Second())
			cond.L.Unlock()
		}(i, cond)
	}
	time.Sleep(time.Second * 1)
	cond.L.Lock()
	// 唤醒一次
	cond.Signal()
	cond.L.Unlock()
	time.Sleep(time.Second * 5)
	cond.L.Lock()
	//唤醒剩下的所有
	cond.Broadcast()
	cond.L.Unlock()
	time.Sleep(time.Second * 1)
}
