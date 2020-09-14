package main

import "fmt"

func main() {

	var (
		i    int = 32
		prt  *int
		pprt **int
	)
	prt = &i
	pprt = &prt
	fmt.Printf("prt的地址：%v -- prt的值：%v", prt, *prt)
	fmt.Printf("\npprt的地址：%v -- pprt的值：%v -- pprt实际指定内存的值：%v", pprt, *pprt, **pprt)

	/*
	 	prt的地址：0xc0000140b8 -- prt的值：32
		pprt的地址：0xc000006028 -- pprt的值：0xc0000140b8 -- pprt实际指定内存的值：32
	*/
}
