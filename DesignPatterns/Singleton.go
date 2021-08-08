package main

import (
        "fmt"
        "sync"
        "time"
)

type Master struct{

}

var single *Master

var once sync.Once

func CreateMaster() *Master{
        if single != nil{
                fmt.Println("Already created")
                return single;
        }else{
                once.Do(func() {
                        single = &Master{}
                        fmt.Println("Creating first time!")
                })
                return single;
        }
}




func main() {
        obj := CreateMaster()
        fmt.Println(obj)
        obj = CreateMaster()
        fmt.Println(obj)
        obj = CreateMaster()
        fmt.Println(obj)
        
}
