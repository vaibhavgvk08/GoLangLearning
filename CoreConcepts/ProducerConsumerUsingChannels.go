package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func producer(mychan chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Produced : ", i)
		mychan <- i
	}
	wg.Done()
}

func consumer(mychan chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Consumed : ", <-mychan)
	}
	wg.Done()
}
func main() {
	mychan := make(chan int, 10)
	wg.Add(1)
	go producer(mychan)
	wg.Add(1)
	go consumer(mychan)

	wg.Wait()
}
