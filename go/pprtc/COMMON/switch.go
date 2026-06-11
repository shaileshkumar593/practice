package main

import (
	"fmt"
)

func main() {
	fmt.Println("Switch case ")
	for i := 0; i < 8; i++ {
		switch i {
		case 0:
			fmt.Println(i)
		case 1:
			fmt.Println(i)
		case 2:
			fmt.Println(i)
		case 3:
			fmt.Println(i)
		default:
			fmt.Println(i)
		}
		fmt.Println("out of for loop")
	}
	fmt.Println("out of main")
}
