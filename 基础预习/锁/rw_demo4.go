package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)
	// 写的时候啥也不能干
	go write1(1)
	go read1(2)
	go write1(3)
	time.Sleep(2 * time.Second)
}
func read1(i int) {
	println(i, "read start")
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	println(i, "read over")
	time.Sleep(1 * time.Second)
}
func write1(i int) {
	println(i, "write start")
	m.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()
	println(i, "write over")
	time.Sleep(1 * time.Second)
}
