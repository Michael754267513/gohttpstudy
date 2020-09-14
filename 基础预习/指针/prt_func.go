package main

import "fmt"

func getsum(a, b *int) (sum int) {

	sum = *a + *b
	return
}

func main() {
	var (
		i = 10
		a = &i
		b = &i
	)
	fmt.Println(getsum(a, b))
}
