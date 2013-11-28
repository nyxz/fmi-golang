package main

type Node struct {
	value int64
	left *Node
	right *Node
}

func (n *Node) WalkInOrder() <- chan int64 {
	var iter func(*Node)
	c := make(chan int64)
	iter := func(n *Node) {
		if n.left != nil {
			iter(n.left)
		}
		c <- n.value
		if n.right != nil {
			iter(n.right)
		}
	}
	go func() {
		iter(n)
		close(c)
	}()
	return c
}

func main() {
	n := make(Node)
	n.value = 1
	n.left := new(Node)
	n.left.value = 2
	n.right := new(Node)
	n.right.value = 3
}