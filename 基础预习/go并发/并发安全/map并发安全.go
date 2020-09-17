package main

import (
	"fmt"
	"sync"
	"time"
)

//var Smap sync.Map

func main() {

	var Smap sync.Map
	go func() { //开一个goroutine写map
		for j := 0; j < 10000; j++ {
			Smap.Store(fmt.Sprintf("%d", j), j)
		}
	}()

	time.Sleep(time.Second * 3)
	Smap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

}
