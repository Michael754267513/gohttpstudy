package main

import "fmt"

func main() {

	for i := 0; i <= 20; i++ {
		if i == 2 {
			continue // 跳过本次循环
		}
		if i == 4 {
			goto GOTO // 跳出循环 也会终止循环
		}
		if i == 6 {
			break // 终止 循环
		}
		fmt.Println(i)
	}
GOTO:
	fmt.Println("goto ")
}
