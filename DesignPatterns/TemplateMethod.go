package main

import (
	"fmt"
)

type iSendMessage interface{
	send_on_interface()
} 

type Nodes struct{ 
	sender iSendMessage
}

func (n *Nodes)SendMessage(){	// like algorithm
	fmt.Println("Prepare")
	n.sender.send_on_interface()
	fmt.Println("Exit")
}

type NodeABC struct{
	Nodes	// just to show inheritance/Composition may be ?
}

func (n *NodeABC)send_on_interface(){	// like algorithm
	fmt.Println("NodeABC sent")
}

type NodeXYZ struct{
	Nodes	// just to show inheritance/Composition may be ?
}

func (n *NodeXYZ)send_on_interface(){	// like algorithm
	fmt.Println("NodeXYZ sent")
}

func main() {
	abc := &NodeABC{}
	xyz := &NodeXYZ{}
	
	Node := &Nodes{
		sender : abc,
	}
	Node.SendMessage()
	
	Node.sender = xyz
	Node.SendMessage()
}
