// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type sample struct{
	a int
}

func passbyvalue(s sample){
	s.a = 10
}

func passbyaddr(s *sample){
	s.a = 10
}

func main() {
	s := sample{}
	fmt.Println(s.a)
	
	passbyvalue(s)
	fmt.Println(s.a)
	
	passbyaddr(&s)
	fmt.Println(s.a)
	
}
