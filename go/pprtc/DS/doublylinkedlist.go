package main

import (
	"fmt"
	"os"
)

type DLL struct {
	left  *DLL
	val   int
	right *DLL
}

var rootdll *DLL = nil
var endDll *DLL = nil

func InsertFrontDll(temp *DLL) {
	if rootdll == nil {
		rootdll = temp
	} else {
		temp.right = rootdll
		rootdll.left = temp
		endDll = rootdll
		rootdll = temp
	}
}

func DisplayDll() {
	fmt.Print("DLL: ")
	if rootdll == nil {
		fmt.Println("Empty list")
	} else {
		var cur *DLL = rootdll

		for cur != nil {
			fmt.Printf("%d -------------->", cur.val)
			cur = cur.right
		}

	}

}

func main() {
	for {
		fmt.Println("1. InsertFront 2.Display 3 InsertDllEnd 4. DeleteFront 5.DeleteEnd 6.InsertPos 7.Exit")
		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter node value: ")
			var ele int
			fmt.Scan(&ele)

			temp := new(DLL)
			temp.val = ele
			temp.left = nil
			temp.right = nil

			InsertFrontDll(temp)

		case 2:
			fmt.Println("Displaying node ")
			DisplayDll()

		case 3:
			InsertDllEnd()
		case 4:
			DeleteFront()
		case 5:
			DeleteEnd()
		case 6:
			InsertPos()
		case 7:
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice")
		}
	}
}
