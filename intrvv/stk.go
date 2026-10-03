package main

import "fmt"

const size = 5

var stack []int

// len(slice) >= 0 not check for -ve
// Push adds an element to the stack
func Push(ele int) {
	if len(stack) >= size {
		fmt.Println("Stack Overflow")
		return
	}

	stack = append(stack, ele)

	fmt.Println("Element pushed:", ele)
}

// Pop removes the top element
func Pop() {
	if len(stack) == 0 {
		fmt.Println("Stack Underflow")
		return
	}

	ele := stack[len(stack)-1]

	stack = stack[:len(stack)-1]

	fmt.Println("Element popped:", ele)
}

// Peek displays the top element
func Peek() {
	if len(stack) == 0 {
		fmt.Println("Stack is empty")
		return
	}

	fmt.Println("Top element:", stack[len(stack)-1])
}

// Display displays all available elements
func Display() {
	if len(stack) == 0 {
		fmt.Println("Stack is empty")
		return
	}

	fmt.Println("Stack:", stack)
}

func main() {

	for {

		fmt.Println("\n========== STACK ==========")
		fmt.Println("1. Push")
		fmt.Println("2. Pop")
		fmt.Println("3. Peek")
		fmt.Println("4. Display")
		fmt.Println("5. Exit")
		fmt.Println("===========================")

		var choice int

		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var ele int

			fmt.Print("Enter element: ")
			fmt.Scan(&ele)

			Push(ele)

		case 2:
			Pop()

		case 3:
			Peek()

		case 4:
			Display()

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
