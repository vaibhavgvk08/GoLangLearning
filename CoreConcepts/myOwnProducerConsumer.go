package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var slice []int
var wg sync.WaitGroup

func criticalSection(i int) bool {
	mutex.Lock()
	if i == 0 {
		//consumer
		if len(slice) > 0 {
			fmt.Println("Consumed : ", slice[0])
			slice = slice[1:]
		} else {
			fmt.Println("No elements")
			mutex.Unlock()
			return false
		}
	} else {
		//producer
		fmt.Println("Produced : ", i)
		slice = append(slice, i)
	}
	mutex.Unlock()
	return true
}

func producer() {
	for i := 1; i <= 10; i++ {
		criticalSection(i)
	}
	wg.Done()
}

func consumer() {
	eatAll := 10
	for eatAll != 0 {
		t := criticalSection(0)
		if t {
			eatAll--
		}
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go producer()
	wg.Add(1)
	go consumer()

	wg.Wait()
}
