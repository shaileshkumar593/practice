package main

import (
	"fmt"
	"os"
)

const SIZE = 8

// Fixed array
var stack [SIZE]int
var top = -1

// Push operation
func push(ele int) {
	if top >= SIZE-1 {
		fmt.Println("Stack Overflow")
		return
	}

	top++
	stack[top] = ele
	fmt.Println("Pushed:", ele, "| Top at:", top)
}

// Pop operation
func pop() int {
	if top == -1 {
		fmt.Println("Stack Underflow")
		return 0
	}

	deleted := stack[top]
	top--
	return deleted
}

// Display stack
func display() {
	if top == -1 {
		fmt.Println("Stack is empty")
		return
	}

	fmt.Println("Index  ----------  Value")
	for i := 0; i <= top; i++ {
		fmt.Println(i, " ---------- ", stack[i])
	}
}

func main() {
	var choice int
	var ele int

	for {
		fmt.Println("\nEnter choice: 1.Push  2.Pop  3.Display  4.Exit")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		switch choice {

		case 1:
			fmt.Print("Enter element to push: ")
			_, err := fmt.Scan(&ele)
			if err != nil {
				fmt.Println("Invalid element")
				continue
			}
			push(ele)

		case 2:
			value := pop()
			if top != -1 {
				fmt.Println("Deleted element is:", value)
			}

		case 3:
			display()

		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice")
		}
	}
}
