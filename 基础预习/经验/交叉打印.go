package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	for i := 0; i <= 24; i++ {
		go func(i int) {
			fmt.Printf("%v", i)
		}(i)
	}
	s := "asdfghjkkllzxcvbnm"
	for _, str1 := range strings.Split(s, "") {
		go func(str1 string) {
			fmt.Printf("%v", str1)
		}(str1)
	}
	time.Sleep(time.Second)
}
