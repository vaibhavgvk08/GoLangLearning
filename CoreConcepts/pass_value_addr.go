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

type Chrome struct{
	myplugin plugin
}

func (c Chrome)surf_internet() {
	fmt.Println("Chrome Suuring!")
	c.myplugin.using_plugin()
} 

func (c Chrome)install_plugin(p plugin) {
	c.myplugin = p
	fmt.Println(p)
} 

func (c *Chrome)install_new_plugin(p plugin) {
	c.myplugin = p
	fmt.Println(p)
} 

type plugin interface{
	using_plugin()
}

type add_blocker_plugin struct{
	a int
}

func (c add_blocker_plugin) using_plugin() {
	fmt.Println("Add Blocked!")
} 


func main() {
	// Simple example
	s := sample{}
	fmt.Println(s.a)
	
	passbyvalue(s)
	fmt.Println(s.a)
	
	passbyaddr(&s)
	fmt.Println(s.a)
	
	// Example using methods
	chrome := Chrome{}
	i := add_blocker_plugin{}
	chrome.install_plugin(i)
	fmt.Println(chrome)
	//chrome.surf_internet() ---> throws error because method has func (c Chrome)install_plugin(p plugin) this signature,  reciever is pass by value. Mutation on object wont be there after we exit that method.
	
	chrome.install_new_plugin(i)
	fmt.Println(chrome)
	chrome.surf_internet()  ---> Works becuase reciever is pass by addr, we have mutated the object and its still persistent 
}
