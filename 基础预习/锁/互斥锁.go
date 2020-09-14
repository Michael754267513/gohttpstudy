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
	var test string
	fmt.Println("start lock main")
	mutex.Lock()
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Println("start lock ", i)
			mutex.Lock()
			test = fmt.Sprintf("test %v", i)
			fmt.Println("get locked ", i)
			fmt.Printf(test)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlock the lock main")
	mutex.Unlock()
	time.Sleep(time.Second)
	/*
		start lock main
		start lock  3
		start lock  1
		start lock  2
		Unlock the lock main
		get locked  3 //解锁后只有一个人会得到

	*/
}
