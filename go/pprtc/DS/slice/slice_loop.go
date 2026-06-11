package main

import "fmt"

func main() {
	arry := []int{6, 9, 3, 5, 22, 76, 86}

	// Iterate using range
	for i := range arry {
		fmt.Println(arry[i])
	}

	fmt.Println("************* other way ********************")
	// Alternative: iterate using value directly
	for _, val := range arry {
		fmt.Println(val)
	}
}
