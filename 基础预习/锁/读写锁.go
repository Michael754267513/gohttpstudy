package main

import (
	"fmt"
	"sync"
	"time"
)

func wirte(i int) (num int) {
	num = 10
	num = i + num

	return num
}

/*
	简单点理解，我不死你永远得不到财产，
	读写锁，写优先，其他读的被锁定，当写锁解锁，就可以执行读操作
	读写锁的访问控制规则如下：
		①多个写操作之间是互斥的
		②写操作与读操作之间也是互斥的
		③多个读操作之间不是互斥的
	在这样的控制规则下，读写锁可以大大降低性能损耗
*/
func main() {
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			rwm.RLock()
			fmt.Println("get locked ", i)
			i = wirte(i)
			time.Sleep(time.Second * 2)
			fmt.Println("try to unlock for reading ", i)
			rwm.RUnlock()
			fmt.Println("unlocked for reading ", i)
		}(i)
	}
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("try to lock for writing")
	rwm.Lock()
	fmt.Println("locked for writing")
}
