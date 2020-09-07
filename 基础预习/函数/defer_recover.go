package main

import "fmt"

func ErrorRecover() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("GG")
		}
	}()
	panic("err")
}

func main() {

	ErrorRecover()
}
