package main

import (
	"fmt"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// insertLevelOrder inserts values into the tree level by level
func insertLevelOrder(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
			i++
		}
		if i < len(values) {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
			i++
		}
	}
	return root
}

// printLevelByLevel prints the tree level by level
func printLevelByLevel(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				fmt.Printf("%d ", node.Val)
				queue = append(queue, node.Left, node.Right)
			} else {
				fmt.Printf("N ") // N represents nil/empty node
			}
		}
		fmt.Println()
	}
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := insertLevelOrder(values)

	fmt.Println("Binary Tree Level by Level:")
	printLevelByLevel(root)
}
