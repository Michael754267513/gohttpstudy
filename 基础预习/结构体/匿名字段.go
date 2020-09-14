package main

import "fmt"

type Book struct {
	string
	int
}

func main() {
	var b Book
	b.int = 10
	b.string = "哈哈"
	fmt.Println(b)
	//{哈哈 10}

}
