package main

import "fmt"

func main() {
	stack := []int{}

	for {
		fmt.Println("\n1. Push")
		fmt.Println("2. Pop")
		fmt.Println("3. Peek")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var val int
			fmt.Print("Enter value: ")
			fmt.Scan(&val)

			stack = append(stack, val)
			fmt.Println("Stack:", stack)

		case 2:
			if len(stack) == 0 {
				fmt.Println("Stack is empty")
			} else {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Println("Popped:", top)
				fmt.Println("Stack:", stack)
			}

		case 3:
			if len(stack) == 0 {
				fmt.Println("Stack is empty")
			} else {
				fmt.Println("Top:", stack[len(stack)-1])
			}

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}