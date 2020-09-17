package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type LockTest interface {
	Inc()
	Load() int32
	Who() string
}

// 互斥锁
type TestMutex struct {
	sum  int32
	lock sync.Mutex
}

func (t *TestMutex) Inc() {
	t.lock.Lock()
	t.sum += 1
	t.lock.Unlock()

}

func (t *TestMutex) Load() int32 {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.sum
}

func (t *TestMutex) Who() string {
	return "互斥锁"
}

// 原子操作
type TestAtomic struct {
	sum int32
}

func (t *TestAtomic) Inc() {
	atomic.AddInt32(&t.sum, +1)
}

func (t *TestAtomic) Load() int32 {
	return atomic.LoadInt32(&t.sum)
}

func (t *TestAtomic) Who() string {
	return "原子操作"
}

func test(LockTest LockTest) {
	start := time.Now()
	var ws sync.WaitGroup
	for i := 0; i < 100000; i++ {
		ws.Add(1)
		go func() {
			LockTest.Inc()
			ws.Done()
		}()
	}
	ws.Wait()
	end := time.Now()
	fmt.Printf("\n值：%v ---操作类型： %v --- 耗时：%v", LockTest.Load(), LockTest.Who(), end.Sub(start))

}

func main() {

	MutexTest := TestMutex{}
	AtomicTest := TestAtomic{}
	test(&MutexTest)
	test(&AtomicTest)
	/*
		值：100000 ---操作类型： 互斥锁 --- 耗时：34ms
		值：100000 ---操作类型： 原子操作 --- 耗时：26ms
	*/
}
