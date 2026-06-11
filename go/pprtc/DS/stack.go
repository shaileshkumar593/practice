package main

import (
	"fmt"
	"os"
)

const SIZE = 8

var stack = make([]int, 0, SIZE) // dynamic slice with capacity
var top = -1                     // top index

func push(ele int) {
	if top >= SIZE-1 {
		fmt.Println("Stack Overflow")
		return
	}
	top++
	if top < len(stack) {
		stack[top] = ele
	} else {
		stack = append(stack, ele)
	}
	fmt.Println("Pushed:", ele, " | Top at:", top)
}

func pop() int {
	if top == -1 {
		fmt.Println("Stack Underflow")
		return 0
	}
	delele := stack[top]
	// Remove top element
	stack = stack[:top]
	top--
	return delele
}

func Display() {
	if top == -1 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Index", " ---------- ", "Value")
	for i := 0; i <= top; i++ {
		fmt.Println(i, " ---------- ", stack[i])
	}
}

func main() {
	var choice int
	var ele int

	for {
		fmt.Println("Enter choice: 1. push  2. pop  3. display  4. Exit")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Element to push to stack: ")
			_, err := fmt.Scan(&ele)
			if err != nil {
				fmt.Println("Invalid element:", err)
				continue
			}
			push(ele)

		case 2:
			fmt.Println("Deleted element is", pop())

		case 3:
			Display()

		case 4:
			fmt.Println("Exiting program...")
			os.Exit(0)

		default:
			fmt.Println("Enter correct choice")
		}
	}
}
