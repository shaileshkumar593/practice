package main

import "fmt"

// createMatrix creates a matrix with given rows and columns,
// and fills it with incremental values.
func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	val := 1
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = val
			val++
		}
	}
	return matrix
}

// printMatrix prints a matrix in row × column format
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%3d ", col) // aligned formatting
		}
		fmt.Println()
	}
}

func main() {
	rows, cols := 4, 5

	// Create matrix
	matrix := createMatrix(rows, cols)

	// Print matrix
	fmt.Println("Matrix:")
	printMatrix(matrix)

	// Example: access and modify an element
	matrix[2][3] = 99
	fmt.Println("\nAfter modification:")
	printMatrix(matrix)
}
