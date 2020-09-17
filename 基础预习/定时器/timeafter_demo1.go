package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.After(time.Second * 3)
	fmt.Printf("t type=%T\n", t)
	//阻塞3秒
	fmt.Println("t=", <-t)
}
