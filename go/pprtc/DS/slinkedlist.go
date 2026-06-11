package main

import "fmt"

type DLL struct {
	left  *DLL
	val   int
	right *DLL
}

var root *DLL = nil

func InsertFront(temp *DLL) {
	if root == nil {
		root = temp
	} else {
		temp.right = root
		root.left = temp
		root = temp
	}
}

func Display() {
	fmt.Print("DLL: ")
	for node := root; node != nil; node = node.right {
		fmt.Printf("%d ", node.val)
	}
	fmt.Println()
}

func main() {
	for {
		fmt.Println("1. InsertFront 2.Display 3.Exit")
		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice) // only accept var

		switch choice {
		case 1:
			fmt.Print("Enter node value: ")
			var ele int
			fmt.Scan(&ele)

			temp := new(DLL) // allocate and return value
			temp.val = ele
			temp.left = nil
			temp.right = nil

			InsertFront(temp)

		case 2:
			Display()

		case 3:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
