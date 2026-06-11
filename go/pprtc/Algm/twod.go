package main

import "fmt"

func main() {
	// ...

	numRows := 4
	board := make([][]bool, numRows)

	// make all the rows of board
	for row := range board {
		fmt.Println("  row ", row)
		board[row] = make([]bool, row)
	}

	// appending false to the end of each row
	for row := range board {
		board[row] = append(board[row], false)
	}

	// append a new row to the board
	newRow := make([]bool, 5)
	board = append(board, newRow)

	fmt.Println(board)
}
