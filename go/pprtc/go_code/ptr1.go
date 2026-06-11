package main

import "fmt"

func main() {
	ptr := new(int)

	fmt.Println("Before change ptr", *ptr)
	// ptr is initialize to nil value 0(zero)
	changePtr(ptr)
	fmt.Println("After Change ptr", *ptr)
}

func changePtr(ptr *int) {
	*ptr = 20
}
