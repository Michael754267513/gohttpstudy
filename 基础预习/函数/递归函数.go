package main

import "fmt"

func recursive(i int) (v int) {

	fmt.Println(i)
	if i == 1 {
		return 1
	}
	return recursive(i - 1)
}

func main() {

	recursive(3)

	/*
		3
		2
		1
	*/

}
