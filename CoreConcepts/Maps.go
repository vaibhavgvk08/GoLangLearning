package main

import (
	"fmt"
)

func myfunc(i map[string]string){
	i["teacher"] = "student"
	fmt.Println(i)
}

func main() {
	mymap := make(map[string]string)
	
	mymap["vaibhav"] = "Engineer"
	mymap["someone"] = "something"
	
	fmt.Println(mymap)
	myfunc(mymap)
	fmt.Println(mymap)
	
	delete(mymap, "someone")
	fmt.Println(mymap)
	
	k, v := mymap["vaibhav"]
	fmt.Println(k, v)
	k, v = mymap["kumar"]
	fmt.Println(k, v)
}
