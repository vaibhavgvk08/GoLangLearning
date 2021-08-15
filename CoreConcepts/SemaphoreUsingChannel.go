package main

import (
	"fmt"
	"time"
)


func criticalSection(semaphore chan int, i int){
	semaphore <- i //somedata
	
	fmt.Println("Accessing Done by : ", i)
}

func main() {
	semaphore := make(chan int, 10)
	
	for i:=0;i<20;i++{
		go criticalSection(semaphore, i)
		<- semaphore //  reset one process done!
	}
	
	time.Sleep(2)
}
