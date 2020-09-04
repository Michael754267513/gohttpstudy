package main

import "fmt"

func main() {
	grow1 := make([]int32, 1, 2)
	grow1 = append(grow1, 1, 2, 3, 4, 5, 6, 7, 8, 8)
	fmt.Println(grow1)

}
