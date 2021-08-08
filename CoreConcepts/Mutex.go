package main

import (
        "fmt"
        "sync"
)

var a int

var wg sync.WaitGroup    // <------- Required to wait till all go routines are done

var mut sync.Mutex

func fun(i int) {
        fmt.Println("i = ", i)
        mut.Lock()       // <------------- Every go routine will reach till here and wait till lock is acquired
        defer mut.Unlock()
        if a%2 == 0{
                a++
        }else{
                a--
        }
        wg.Done()
}

func main() {

        for i:=0;i<10;i++{
                wg.Add(1)
                go fun(i)
        }
        wg.Wait()

        fmt.Println(a)
}
