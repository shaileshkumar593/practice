package main

import "fmt"

func picTwo(a []int) {
	var firstLargest, secondLargest int

	for _, val := range a {
		if val > firstLargest {
			secondLargest = firstLargest
			firstLargest = val
		} else if val > secondLargest {
			secondLargest = val
		}
	}

	fmt.Println("First Largest:", firstLargest)
	fmt.Println("Second Largest:", secondLargest)
}

func main() {
	s := []int{7, 2, 5, 6, 3}

	picTwo(s)
}
