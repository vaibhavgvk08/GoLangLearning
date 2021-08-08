package main

import (
	"fmt"
)

type iMaster interface{
	Teach()
}

type danceMaster struct{	
}

func (d danceMaster) Teach(){
	fmt.Println("Teach Dance")
}

type sportMaster struct{	
}

func (d sportMaster) Teach(){
	fmt.Println("Teach Sports")
}

//Factory that creates masters
func CreateMaster(s string) iMaster{
	if s == "Dance"{
		return danceMaster{}
	}else{
		return sportMaster{}
	}
}


func main(){
	var m iMaster
	
	m = CreateMaster("Dance")
	m.Teach()
	
	m = CreateMaster("Sports")
	m.Teach()

}
