package main

import "fmt"

func main() {

	var slice []int          //initialize a slice
	slice = append(slice, 1) //fill the slice using append method
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)

	fmt.Println("The slice created by user is:", slice)

	if slice == nil {
		fmt.Println("Slice is empty")
	} else {
		fmt.Println("Slice is not empty")
	}
}
