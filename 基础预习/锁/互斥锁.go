package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	// 多人竞争 只有一个人活得

	简单理解我死后 只有一个人可以获得遗产

	老子有一块蛋糕，你们居然都想吃，你们自己竞争下，最后只有一个人可以获得老子蛋糕
*/
func main() {
	var mutex sync.Mutex
	fmt.Println("start lock main")
	mutex.Lock()
	fmt.Println("get locked main")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Println("start lock ", i)
			mutex.Lock()
			fmt.Println("get locked ", i)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlock the lock main")
	mutex.Unlock()
	fmt.Println("get unlocked main")
	time.Sleep(time.Second)
}
