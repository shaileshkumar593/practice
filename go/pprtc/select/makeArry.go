package main

import "fmt"

func main() {
	// Using make() to create a slice of integers
	numbers := make([]int, 5, 10)

	// Displaying the slice's length, capacity, and values
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
	fmt.Println("Values:", numbers)

	// Using make() to create a slice of integers
	numbersWithoutOptional := make([]int, 5)

	// Displaying the slice's length, capacity, and values
	fmt.Println("Length:", len(numbersWithoutOptional))
	fmt.Println("Capacity:", cap(numbersWithoutOptional))
	fmt.Println("Values:", numbersWithoutOptional)
}
