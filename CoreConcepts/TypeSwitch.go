package main

import "fmt"

type circle struct {
}

func (c circle) drawcircle() {
	fmt.Println("circle")
}

func (c circle) draw1() string {
	return "circle 1"
}

type square struct {
}

func (c square) drawsquare() {
	fmt.Println("square")
}

func (c square) draw1() string {
	return "square 1"
}

func doer(i interface{}) {
	switch v := i.(type) {
	case circle:
		v.drawcircle()
	case square:
		v.drawsquare()
	}
}

func main() {
	c := circle{}
	sq := square{}
	doer(c)
	doer(sq)
}
