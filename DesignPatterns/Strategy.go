// You can edit this code!
// Click here and start typing.
package main

import "fmt"


type Numbers struct{
	s Sort 
}

func (n *Numbers)SortNumbers(){
	fmt.Print("Sorting using ...")
	n.s.sort()
}

func (n *Numbers)SetStrategy(so Sort){
	n.s = so
}

type Sort interface{
	sort()
}

type QuickSort struct{}

func (n QuickSort)sort(){
	fmt.Println("Quick Sort")
}

type MergeSort struct{}

func (n MergeSort)sort(){
	fmt.Println("Merge Sort")
}

func main() {
	n := Numbers{}
	q := QuickSort{}
	n.SetStrategy(q)
	n.SortNumbers()
	
	m := MergeSort{}
	n.SetStrategy(m)
	n.SortNumbers()	
	
}
