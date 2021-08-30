package main

import "fmt"

func fun(c, quit chan int) {
	
	for {
		select {
		case i := <- c:
			fmt.Println(i)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <-i
		}
		quit <- 0
	}()
	fun(c, quit)
}
