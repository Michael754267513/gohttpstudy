package main

import (
	"fmt"
	"sync"
)

var ws sync.WaitGroup

func Send(jobChan chan int) {
	for a := 0; a <= 100; a++ {
		ws.Add(1)
		jobChan <- a
	}
	close(jobChan)
}

func SumCH(jobChan chan int, resultChan chan int) {
	var sum int
	for {
		i, ok := <-jobChan
		if !ok {
			break
		}
		sum += i
		resultChan <- sum
	}

}

func GetSum(resultChan chan int) {
	for {
		i, ok := <-resultChan
		if !ok {
			break
		}
		fmt.Println(i)
		ws.Done()
	}
	close(resultChan)
}

func main() {
	jobChan := make(chan int, 10000)
	resultChan := make(chan int, 10000)
	go Send(jobChan)
	for i := 0; i < 24; i++ {
		go SumCH(jobChan, resultChan)
	}
	go GetSum(resultChan)
	ws.Wait()
}
