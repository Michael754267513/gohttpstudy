package main

import (
	"fmt"
	"strings"
)

func SortStr(s string) (sort_s string) {
	sliceStr := strings.Split(s, "")
	for i := len(sliceStr) - 1; i >= 0; i-- {
		sort_s += sliceStr[i]
	}

	return sort_s
}

func main() {
	fmt.Printf(SortStr("asdfgh"))
}
