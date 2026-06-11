package main

import "fmt"

func main() {
	// Number of rows
	rows := 4

	// Create ragged matrix
	ragged := make([][]int, rows)

	// Each row has different number of columns
	for i := 0; i < rows; i++ {
		ragged[i] = make([]int, i+1) // row length grows with i
		for j := 0; j < len(ragged[i]); j++ {
			ragged[i][j] = (i+1)*10 + j // fill with values
		}
	}

	// Print ragged matrix
	fmt.Println("Ragged Matrix:")
	for _, row := range ragged {
		fmt.Println(row)
	}
}
