package main

import (
	"fmt"
)

// Node structure
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert into BST
func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}

	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else {
		root.Right = insert(root.Right, value)
	}

	return root
}

// Inorder Traversal (Left → Root → Right)
func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Value, " ")
	inorder(root.Right)
}

// Preorder Traversal (Root → Left → Right)
func preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Value, " ")
	preorder(root.Left)
	preorder(root.Right)
}

// Postorder Traversal (Left → Right → Root)
func postorder(root *Node) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Print(root.Value, " ")
}

// Level Order Traversal (BFS)
func levelOrder(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Print(current.Value, " ")

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

// Search in BST
func search(root *Node, key int) bool {
	if root == nil {
		return false
	}

	if root.Value == key {
		return true
	}

	if key < root.Value {
		return search(root.Left, key)
	}

	return search(root.Right, key)
}

// Height of Tree
func height(root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func main() {
	var root *Node

	values := []int{10, 5, 20, 3, 7, 30}

	// Insert values
	for _, v := range values {
		root = insert(root, v)
	}

	fmt.Println("Inorder Traversal:")
	inorder(root)

	fmt.Println("\nPreorder Traversal:")
	preorder(root)

	fmt.Println("\nPostorder Traversal:")
	postorder(root)

	fmt.Println("\nLevel Order Traversal:")
	levelOrder(root)

	fmt.Println("\nSearch 7:", search(root, 7))
	fmt.Println("Search 100:", search(root, 100))

	fmt.Println("Height of Tree:", height(root))
}
