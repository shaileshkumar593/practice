package main

import (
	"fmt"
)

type Node struct {
	x    int
	next *Node
}

func main() {
	var s *Node = nil
	s = new(Node)
	s.x = 20
	s.next = nil
	var p Node
	p.x = 25
	p.next = nil

	fmt.Println(s, &s)
	fmt.Println(p, &p)
	fmt.Printf("%p\n", &p)
	fmt.Printf("%p\n", s)

}
