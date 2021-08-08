package main

import (
	"fmt"
)

type computer interface{
	connectTypeC()
}

type mac struct{}

func (m mac) connectTypeC() {
	fmt.Println("Typec Connected")
}

type windows struct{}

func (w windows) connectPin() {
	fmt.Println("pin connected")
}

//------
type windowsAdapter struct{
	w windows
}

func (win windowsAdapter) connectTypeC() {
	win.w.connectPin()
	fmt.Println("connectTypeC connected")
}


func main() {
    var m = mac{}
    m.connectTypeC()
    
    var win windows = windows{}
    winada := windowsAdapter{
	w : win,
    }
    winada.connectTypeC()
    
}
