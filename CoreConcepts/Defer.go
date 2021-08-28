package main

import "fmt"

func CallMeLast() {
	fmt.Println("Last called !")
}

func returnThis() (int, error) {
	i, err := fmt.Println("third called")
	return i, err
}

func test() (n int, err error) {
	fmt.Println("First called!")
	defer CallMeLast()
	fmt.Println("Second called!")
	return returnThis()
}

func main() {
	test()
}
