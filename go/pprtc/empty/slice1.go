package main

import "fmt"

//create a function main
func main() {

	var slice []int // initialize slice

	slice = append(slice, 1) //fill the slice using append function
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)

	fmt.Println("The slice created by user is:", slice)

	if len(slice) == 0 {
		fmt.Println("Slice is empty")
	} else {
		fmt.Println("Slice is not empty")
	}

}
