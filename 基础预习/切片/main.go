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
	makecheck := make([]int32, 10, 20) // 长度， 容量
	fmt.Println(newcheck)              // new 返回的是一个指针对象 new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值
	fmt.Println(*newcheck)             //
	fmt.Println(makecheck)             // make只能创建slice、map和channel类型的数据，并且返回一个有初始值(非零)的T类型，而不是*T

	fmt.Println("切片容量:", cap(makecheck)) // 获取切片的容量
	fmt.Println("切片长度：", len(makecheck)) // 获取切片实际使用长度
	/*
		切片容量: 20
		切片长度： 10
	*/

	makecheck = append(makecheck, 1, 2, 3, 4, 5, 6, 7) // 切片赋值
	fmt.Println("切片容量:", cap(makecheck))               // 获取切片的容量
	fmt.Println("切片长度：", len(makecheck))               // 获取切片实际使用长度
	fmt.Println("切片值：", makecheck)
	/*
		切片容量: 20
		切片长度： 17
		切片值： [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7]
	*/
	s1 := [2]int32{1}
	s2 := []int32{}
	s3 := [2]int32{}
	fmt.Println("切片s1:", s1)
	fmt.Println("切片s1长度:", len(s1))
	fmt.Println("切片s1容量:", cap(s1))
	fmt.Println("切片s2:", s2)
	fmt.Println("切片s2长度:", len(s2))
	fmt.Println("切片s2容量:", cap(s2))
	fmt.Println("切片s3:", s3)
	fmt.Println("切片s3长度:", len(s3))
	fmt.Println("切片s3容量:", cap(s3))
	/*
		1. 申明切片长度 如果没有值，会按照默认值填充
		2. 不申明切片长度,默认容量和长度都为0，为空切片
		切片s1: [1 0]
		切片s1长度: 2
		切片s1容量: 2
		切片s2: []
		切片s2长度: 0
		切片s2容量: 0
		切片s3: [0 0]
		切片s3长度: 2
		切片s3容量: 2
	*/
	scopy1 := []int32{1, 2, 3, 4, 5}
	scopy2 := []int32{1, 1, 1, 1}
	scopy3 := make([]int32, 1, 10)
	scopy4 := []int32{}
	copy(scopy2, scopy1)
	copy(scopy3, scopy1)
	copy(scopy4, scopy1)
	copy(scopy4, scopy1)
	fmt.Println("scopy1: ", scopy1)
	fmt.Println("scopy2: ", scopy2)
	fmt.Println("scopy3: ", scopy3)
	fmt.Println("scopy4: ", scopy4)
	/*
		# 注意拷贝目的切片的长度，如果没有大于或者等于源切片长度name 会损失部分值
		scopy1:  [1 2 3 4 5]
		scopy2:  [1 2 3 4]
		scopy3:  [1]
	*/
	scopy2[1] = 10
	fmt.Println("修改scopy2 值: ", scopy2) // [1 10 3 4 ]
	// 切片扩容
	s := make([]int, 1, 3)
	s = append(s, 1, 5, 6, 7, 8, 1, 1, 1)
	fmt.Println("切片s:", s)
	fmt.Println("切片s长度:", len(s))
	fmt.Println("切片s容量:", cap(s))
	/*
		# 源码slice的长度在超过一个阈值（这里是1024）后便不再翻倍，而是每次以25%的幅度增长，直到满足所需的容量
		实际切片扩容 与切片数据类型 原本切片容量 和容量空间都有关系
		GoLang中的切片扩容机制，与切片的数据类型、原本切片的容量、所需要的容量都有关系，比较复杂。对于常见数据类型，在元素数量较少时，大致可以认为扩容是按照翻倍进行的。
		为了避免因为切片是否发生扩容的问题导致bug，最好的处理办法还是在必要时使用 copy 来复制数据，保证得到一个新的切片，以避免后续操作带来预料之外的副作用。
	*/
}
