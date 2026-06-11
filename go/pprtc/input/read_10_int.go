package main

import (
	"fmt"
	"os"
)

func readIntSlice(count int) ([]int, error) {
	slice := make([]int, 0, count)
	var num int

	fmt.Printf("Enter %d integers (space-separated or one per line):\n", count)

	for i := 0; i < count; i++ {
		_, err := fmt.Scan(&num)
		if err != nil {
			return nil, fmt.Errorf("failed to read integer at position %d: %w", i+1, err)
		}
		slice = append(slice, num)
	}

	return slice, nil
}

func main() {
	const numElements = 10

	slice, err := readIntSlice(numElements)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("You entered:")
	for i, val := range slice {
		fmt.Printf("Element %d: %d\n", i, val)
	}
}
