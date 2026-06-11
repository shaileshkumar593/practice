package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (root *Node) PreOrderTraversal() {
	if root != nil {
		fmt.Printf("%d ", root.data)
		root.left.PreOrderTraversal()
		root.right.PreOrderTraversal()
	}
	return
}
func main() {
	root := Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = &Node{4, nil, nil}
	root.left.right = &Node{5, nil, nil}
	root.right.left = &Node{6, nil, nil}
	root.right.right = &Node{7, nil, nil}
	fmt.Printf("Pre Order Traversal of the given tree is: ")
	root.PreOrderTraversal()
}
