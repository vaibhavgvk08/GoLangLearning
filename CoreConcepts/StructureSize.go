package main

import (
	"fmt"
	"unsafe"
)

// structures are always multiples of 8 bytes.... padding is used to store structs
type old struct {
	b bool  // 1     8 bytes with padding
	i int64 // 8	8 bytes
	t int32 // 4	8 bytes wuth paddng ==== totoal size is 24
}

func (o old) myfun() { // this does not explaicitly consume memorty in size of structure
	fmt.Print("old fund")
}

type newstrict struct { // efficeient storage
	i int64 // 8	// 8 bytes
	b bool  // 1	// 1 byte and 3 bytes padding
	t int32 // 4	// other part of above bytes ie 4 byets ===== totoal size is 16 only
}

func main() {
	a := old{
		i: 10,
		b: true,
		t: 12,
	}
	fmt.Println(unsafe.Sizeof(a))
	a.myfun()

	b := newstrict{
		i: 10,
		b: true,
		t: 12,
	}
	fmt.Println(unsafe.Sizeof(b))
}
