package main

import "fmt"

type base interface{}

type my struct {
	a int
	b int
}

func (i my) log() string {
	return "logged"
}

type shape interface {
	draw() string
}

type circle struct {
}

func (c circle) draw() string {
	return "circle"
}

func (c circle) draw1() string {
	return "circle 1"
}

type square struct {
}

func (c square) draw() string {
	return "square"
}

func (c square) draw1() string {
	return "square 1"
}

func main() {
	myobj := my{a: 10, b: 20}
	var b base
	b = myobj
	fmt.Println(b.(my).a) // say that b is of type my struct here

	var s shape
	c := circle{}
	sq := square{}

	s = c
	fmt.Println(s.(circle).draw1()) // here we need to say s.(circle).draw1() becuase shape doesnt have draw1
	s = sq
	fmt.Println(s.draw()) // here we dont need to say s.(circle).draw() becuase shape has a draw already
}
