// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// variables
// slices, maps, structs
// functions, methods.

func addVals[T constraints.Ordered](a, b T) T {
	return a + b
}

func returnSquares[T int | float64](a []T) []T {
	var sq []T
	for _, d := range a {
		sq = append(sq, d*d)
	}
	return sq
}

type MyCustomType interface {
	constraints.Ordered
}

type MyStruct[T MyCustomType] struct {
	age T
}

type MyMap[T comparable] map[T]int

type MyMapNew[T comparable, V int | float64] map[T]V

func main() {
	// generic functions.
	fmt.Println("AddVals :: ", addVals(13, 24))
	fmt.Println("AddVals :: ", addVals(1.3, 2.4))

	// generic functions with slices.
	fmt.Println("SquareVals :: ", returnSquares([]int{1, 2, 3}))
	fmt.Println("SquareVals :: ", returnSquares([]float64{1.5, 2.5, 3.5}))

	// generic structs.
	obj := MyStruct[string]{age: "abs"}
	fmt.Println(obj)

	// generic maps
	mym := make(MyMap[string])
	mym["abs"] = 10
	fmt.Println(mym)

	mym1 := make(MyMapNew[string, int])
	mym1["abs"] = 10
	fmt.Println(mym1)

	mym2 := make(MyMapNew[string, float64])
	mym2["abs"] = 10.5
	fmt.Println(mym2)
}
