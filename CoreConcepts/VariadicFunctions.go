package main

import (
	"fmt"
)

func printSlice(something int, arr ...int){ // variadic function -> always needs to be at the end
	fmt.Println(arr)
}


//	func test(i ...int, s ...string){ // ---> wrong usage error
//		fmt.Println(i)
//		fmt.Println(s)
//	}

func test(i ...[]int){ // variadic slice parameters 
	fmt.Println(i)
}

func main() {
	
	slice2 := make([]int, 3) 	//slice of 3 len
	
	slice2 = append(slice2, 1,2,3)
	
	// slice is like a pointer using underlying arrays
	fmt.Println( slice2)
	
	// printSlice(slice2). -> would throw error 
	
	printSlice(0, slice2...) // allow slice to be passed as variadic parameter
	
	var twoD [][]int //2d slice
	
	for i:=0;i<5;i++{
		var newSlice []int
		for j:= 0; j<5;j++{
			newSlice = append(newSlice, j)
		}
		twoD = append(twoD, newSlice)
	} 
	test(twoD...)
}
