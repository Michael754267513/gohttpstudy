package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// 申明一个map   map[KeyType]ValueType
	// map[KeyType]interface{}  value类型可以为任意值，因为空interface可以接受任意值

	str := make(map[int]string)
	str[1] = "测试"
	fmt.Println(str)

	// 任意值
	inter := make(map[string]interface{})
	inter["姓名"] = "Michael"
	inter["age"] = 18
	fmt.Println(inter)

	// 判断值是否存在
	if value, ok := inter["age"]; ok {
		// 存在就是 true
		fmt.Println(value)
	}
	// map 遍历
	for k, v := range inter {
		fmt.Println("key: ", k, "---- value:", v)
	}
	// map 删除
	delete(inter, "age")
	fmt.Println(inter)

	// map 排序 遍历所有的key 存放到切片里面，根据切片进行排序，然后根据排序类容遍历所有的balue
	sort1 := make(map[int]int)
	for i := 0; i < 10; i++ {
		sort1[rand.Intn(100)] = rand.Intn(100)
	}
	key := []int{}
	for k, _ := range sort1 {
		key = append(key, k)
	}
	fmt.Println("排序前：", key)
	sort.Ints(key)
	fmt.Println("排序后：", key)
	fmt.Println("map初始值:", sort1)
	for _, k := range key {
		fmt.Println(sort1[k])
	}

}
