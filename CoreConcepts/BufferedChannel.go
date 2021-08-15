package main

import (
	"fmt"
	"time"
)



func main() {
	mychan := make(chan int, 10)
	
	for i:=1;i<=11;i++{
		mychan <- i //somedata
		if i == 10{
			fmt.Println("r by : ", <-mychan) // to unblock the sending part we are reading one value
		}
	}
	
	fmt.Println("r by : ", <-mychan)
	fmt.Println("r by : ", <-mychan)
	fmt.Println("r by : ", <-mychan)
	fmt.Println("r by : ", <-mychan)
	fmt.Println("r by : ", <-mychan)
	
	fmt.Println("r by : ", <-mychan)
	fmt.Println("r by : ", <-mychan)
	fmt.Println("r by : ", <-mychan)
	fmt.Println("r by : ", <-mychan)
	fmt.Println("r by : ", <-mychan)

	//fmt.Println("r by : ", <-mychan)
		
	time.Sleep(2)
}
