package main

import "fmt"

/*
	优点： 保证数据的安全，节约内存，指针相同的数据共享(当指针的值变化后，所有指针执行该地址的都会发生变化)
*/

func main() {
	// 申明一个指针
	var (
		prt *int32
		i   int32 = 10
	)
	prt = &i                   // 赋值
	fmt.Println("i 的指针地址", &i) // 打印指针地址
	fmt.Print("prt 的指针", prt)  // 打印指针地址
	fmt.Println("\r\n")
	fmt.Println("指针地址的值", *prt) // 获取指针值
	fmt.Print("修改i的值")
	fmt.Println("\r\n")
	i = 50
	fmt.Println("i 的指针地址", &i) // 打印指针地址
	fmt.Print("prt 的指针", prt)  // 打印指针地址
	fmt.Println("\r\n")
	fmt.Println("指针地址的值", *prt) // 获取指针值

	/*
		i 的指针地址 0xc0000140a8
		prt 的指针0xc0000140a8
		指针地址的值 10

		修改i的值
		i 的指针地址 0xc0000140a8
		prt 的指针0xc0000140a8
		指针地址的值 50

	*/

}
