// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type ibrowser interface{
	surf_internet()
	install_plugin(plugin)
}

type Chrome struct{
	myplugin plugin
}

func (c *Chrome)surf_internet() {
	fmt.Println("Chrome Suuring!")
	c.myplugin.using_plugin()
} 

func (c *Chrome)install_plugin(p plugin) {
	c.myplugin = p
	fmt.Println("Chrome PLugin Installed !")
} 

type Firefox struct{
	myplugin plugin
}

func (c *Firefox)surf_internet() {
	fmt.Println("Firefox Suuring!")
	
} 

func (c *Firefox)install_plugin(p plugin) {
	c.myplugin = p
	fmt.Println("Firefox PLugin Installed !")
} 

type plugin interface{
	using_plugin()
}

type add_blocker_plugin struct{
}

func (c add_blocker_plugin) using_plugin() {
	fmt.Println("Add Blocked!")
} 

func main() {
	p := &add_blocker_plugin{}
	chrome := &Chrome{}
	chrome.install_plugin(p)
	chrome.surf_internet()
}
