package main

import "fmt"

func main() {
	/*
			var identifier []type // 申明一个切片
			var slice1 []type = make([]type, len)
			slice1 := make([]type, len)
			s :=[] int {1,2,3 } // 切片初始化

		len() 和 cap() 函数
		切片是可索引的，并且可以由 len() 方法获取长度。
		切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
		切片截取：
			index 由0 开始
	*/

	newcheck := new(int)
	makecheck := make([]int32, 10)
	fmt.Println(newcheck)  // new 返回的是一个指针对象 new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值
	fmt.Println(*newcheck) //
	fmt.Println(makecheck) // make只能创建slice、map和channel类型的数据，并且返回一个有初始值(非零)的T类型，而不是*T
}
