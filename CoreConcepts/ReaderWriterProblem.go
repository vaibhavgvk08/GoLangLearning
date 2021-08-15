// Here we are allowed to do multiple read operations parallelly byt when a write comes it needs to be atomic. Any no of processes are allowed to do read at same time.
package main

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.RWMutex

func reader(i int){
	rw.RLock()
	defer rw.RUnlock()
	
	fmt.Println("Reading Done by : ", i)
}

func writer(i int){
	rw.Lock()
	defer rw.Unlock()
	
	fmt.Println("Writing Done by : ", i)
}

func main() {
	for i:=0;i<10;i++{
		go writer(i)
	}
	for i:=0;i<10;i++{
		go reader(i)
	}
	
	time.Sleep(2)
}
