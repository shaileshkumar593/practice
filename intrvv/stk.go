package main

import (
	"fmt"
)

var top int = -1

func Pushele(s []int, size int, ele int) {
	if len(s) > size {
		fmt.Println("stack overflow")
		return
	}
	top = top + 1
	s[top] = ele

}

func Popele(s []int) (ele int) {
	if len(s) < 0 {
		fmt.Println("stack  underflow")
		return -1
	}
	ele = s[top]
	top = top - 1
	if top < 0 {
		top = 0

	}

	return ele
}

func main() {
	fmt.Println("Hello, World!")
	var size int = 5
	var choice int = 0
	var ele int

	stackList := make([]int, size)

	for {
		fmt.Println("1. push 2. pop ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("enter element to insert")
			fmt.Scan(&ele)
			Pushele(stackList, size, ele)

		case 2:
			fmt.Println(Popele(stackList))
			stackList = stackList[:top]

		default:
			fmt.Println("not correct choice")

		}
	}

}
