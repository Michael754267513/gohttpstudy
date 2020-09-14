package main

import "fmt"

type Computer string

func (C Computer) Action() {
	fmt.Println(C)
}

func main() {
	var dell Computer
	dell = "Dell"
	dell.Action()
}
