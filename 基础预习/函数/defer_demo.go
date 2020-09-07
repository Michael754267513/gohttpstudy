package main

import "fmt"

func main() {
	//非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。
	// 文件打开需要关闭，呢嘛可以在打开的时候 加个defer 关闭的，当处理完毕自动关闭

	// 延迟处理 ，在函数处理完毕，最后处理，先defer的最后处理，
	fmt.Println("Starting")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("End")
}
