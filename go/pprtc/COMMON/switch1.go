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
			break
		case 1:
			fmt.Println(i)
			break
		case 2:
			fmt.Println(i)
			break
		case 3:
			fmt.Println(i)
			break
		default:
			fmt.Println(i)
		}
		fmt.Println("out of for loop")
	}
	fmt.Println("out of main")
}
