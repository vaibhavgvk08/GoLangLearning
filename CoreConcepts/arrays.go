package main

import (
	"fmt"
)

type Master struct{
	arr [2]int 
}

func main() {
	var arr [5]int
	fmt.Println(arr)
	arr1 := [5]int{1,2,3}
	fmt.Println(arr1)
	fmt.Println(arr1[1])
	
	for i:=0;i<3;i++{
		fmt.Println(arr1[i])
	}
	
	for _,v := range arr1 {
		fmt.Println(v)
	}
	
	var arr2 [2]Master
	
	arr2[0] = Master{
		arr : [2]int{1,2},
	}
	arr2[1] = Master{
		arr : [2]int{3,4},
	}
	
	fmt.Println(arr2)
}
