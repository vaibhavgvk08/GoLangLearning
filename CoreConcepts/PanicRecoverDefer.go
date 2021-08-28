package main

import "fmt"

func recur(i int) {
	defer recoverfunc()  // second called,  -> only called once when panic starts other calls wont get any panic so this is not executed
	defer fmt.Println(i) // first called, 3rd, 4th etc....
	if i == 3 {
		err := fmt.Errorf("some error occured")
		panic(err)
	}
	recur(i + 1)
}

func recoverfunc() {
	p := recover()
	if p != nil {
		fmt.Println(p.(error))
	}
}

func main() {
	recur(0)
	fmt.Println("End of main")
}
