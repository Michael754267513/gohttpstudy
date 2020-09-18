package main

import (
	"fmt"
	"time"
)

func main() {
	exitChan := make(chan bool)

	go func() {
		for {
			fmt.Println("Doing Work......")
			select {
			case <-exitChan:
				return
			}
		}
	}()

	time.Sleep(time.Second * 1)
	fmt.Println("stop the goRoutinue")
	close(exitChan)
	time.Sleep(time.Second * 3)
}
