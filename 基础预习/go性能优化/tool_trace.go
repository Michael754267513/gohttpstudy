package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

// 使用标准库runtime/trace  直接运行程序，会生成一个trace.out文件，然后可以进行分析
// go tool trace trace.out 运行 可访问返回的http地址

//var wg sync.WaitGroup
func sleep() {
	time.Sleep(time.Second * 2)
	//wg.Done()
}
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// Your program here
	//for i:=0;i<3;i++ {
	//	//wg.Add(1)
	//	sleep()
	//}
	//wg.Wait()
	fmt.Println("aaaa")
}
