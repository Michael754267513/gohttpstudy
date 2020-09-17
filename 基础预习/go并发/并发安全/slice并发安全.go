package main

import (
	"fmt"
	"sync"
	"time"
)

var ss []int
var lock sync.Mutex

func appendValue1(i int) {
	lock.Lock()
	ss = append(ss, i)
	lock.Unlock()
}

func main() {
	for i := 0; i < 10000; i++ { //10000个协程同时添加切片
		go appendValue1(i)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(len(ss))
}
