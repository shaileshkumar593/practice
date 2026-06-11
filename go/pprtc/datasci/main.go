package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	data := []float64{1.0, 2.0, 3.0, 4.0}
	matrix := mat.NewDense(2, 2, data)

	fmt.Println("Matrix:")
	fmt.Println(mat.Formatted(matrix))

	transposed := mat.NewDense(2, 2, nil)
	transposed.Trace()

	fmt.Println("Transposed Matrix:")
	fmt.Println(mat.Formatted(transposed))
}
