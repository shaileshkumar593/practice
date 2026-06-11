package main

import "fmt"

type StackInt struct {
	s []int
}

// isEmpty() function
func (s *StackInt) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

// length() function
func (s *StackInt) Length() int {
	length := len(s.s)
	return length
}

// Print() function
func (s *StackInt) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

// Push() function
func (s *StackInt) Push(value int) {
	s.s = append(s.s, value)
}

// Pop() function
func (s *StackInt) Pop() int {
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

// Top() function
func (s *StackInt) Top() int {
	length := len(s.s)
	res := s.s[length-1]
	return res
}

func main() {
	var stack StackInt // create a stack variable of type Stack
	/* Adding items to stack */
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Printng the stack:")
	stack.Print() // Print the stack
	fmt.Println("Checking length of stack:")
	fmt.Println(stack.Length()) // Get Length
	fmt.Println("Removing an Item:")
	stack.Pop() // Remove an item
	stack.Print()
	fmt.Println("Getting last Added Item in Stack")
	fmt.Println(stack.Top()) // Get last item
}
