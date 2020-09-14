package main

import (
	"fmt"
	"sync"
)

func param() {
	ws := sync.WaitGroup{}
	s := []string{"a", "b", "c"}
	for _, v := range s {
		ws.Add(1)
		go func(v string) {
			fmt.Printf("\n param: 地址：%v -- param：值：%v", &v, v)
			ws.Done()
		}(v) //每次将变量 v 的拷贝传进函数
	}
	ws.Wait()
}

func Noparam() {
	ws := sync.WaitGroup{}
	s := []string{"a", "b", "c"}
	for _, v := range s {
		ws.Add(1)
		go func() {
			fmt.Printf("\n Noparam: 地址：%v -- Noparam：值：%v", &v, v)
			ws.Done()
		}()
	}
	ws.Wait()
}
func main() {
	param()
	Noparam()
	/*
	 param: 地址：0xc00002c1f0 -- param：值：c
	 param: 地址：0xc00002c210 -- param：值：b
	 param: 地址：0xc000086000 -- param：值：a
	 Noparam: 地址：0xc000086020 -- Noparam：值：c
	 Noparam: 地址：0xc000086020 -- Noparam：值：c
	 Noparam: 地址：0xc000086020 -- Noparam：值：c
	*/
}
