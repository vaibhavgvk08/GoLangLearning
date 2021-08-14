package main

import (
	"fmt"
)

func main() {
	var slice1 []int
	slice2 := make([]int, 3) 	//slice of 3 len
	slice3 := make([]int, 3, 10) 	// slice of 3 len and 10 capacity
	
	// slice is like a pointer using underlying arrays
	fmt.Println(slice1, slice2, slice3)
	
	
	slice4 := slice3
	// pointers 3,4 both pointing to same underlying array
	
	slice4[0] = 1
	fmt.Println(slice3, slice4)
	fmt.Println("Size = ", len(slice3), cap(slice3))
	slice5 := append(slice3, slice4...)
	fmt.Println(slice5) // slice5 may be newly allocated or not its not guarenteed that append will reuse same array or not
	
	var twoD [][]int //2d slice
	
	for i:=0;i<5;i++{
		var newSlice []int
		for j:= 0; j<5;j++{
			newSlice = append(newSlice, j)
		}
		twoD = append(twoD, newSlice)
		fmt.Println(&twoD[i])
	} 
	twoD[0][0] = 10
	fmt.Println(twoD)
}
