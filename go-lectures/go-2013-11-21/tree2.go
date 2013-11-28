package main

import "fmt"

type Node struct {
	value int64
	left *Node
	right *Node
}

func (n *Node) WalkInOrder() <- chan int64 {
	c := make(chan int64)
	
	go func() {
		defer close(c)
		n.iter(c)
	}()
	return c
}

func (n *Node) iter(c chan int64) {
	if n.left != nil {
		n.left.iter(c)
	}
	c <- n.value
	if n.right != nil {
		n.right.iter(c)
	}
}

func main() {
	n := new(Node)
	n.value = 1
	n.left = new(Node)
	n.left.value = 2
	n.right = new(Node)
	n.right.value = 3
	c := n.WalkInOrder()
	for value := range c {
		fmt.Println(value)
	}
}