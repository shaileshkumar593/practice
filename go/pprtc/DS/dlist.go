package main

import "fmt"

type Node struct {
	prev  *Node
	value int
	next  *Node
}

func CreateNewNode(value int) *Node {
	var node Node
	node.next = nil
	node.value = value
	node.prev = nil
	return &node
}
func TraverseDoublyLL(head *Node) {
	// Forward Traversal
	fmt.Printf("Doubly Linked List: ")
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
}
func main() {
	// 10 <=> 20 <=> 30 <=> 40
	head := CreateNewNode(10)
	node_2 := CreateNewNode(20)
	node_3 := CreateNewNode(30)
	node_4 := CreateNewNode(40)
	head.next = node_2
	node_2.prev = head
	node_2.next = node_3
	node_3.prev = node_2
	node_3.next = node_4
	node_4.prev = node_3
	TraverseDoublyLL(head)
}
