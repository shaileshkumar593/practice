package main

/*
	6. ZigZag Conversion
	Problem:
		Convert a string to a zigzag pattern on a given number of rows and read line by line.
*/

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([][]byte, numRows)
	curRow, goingDown := 0, false

	for i := 0; i < len(s); i++ {
		rows[curRow] = append(rows[curRow], s[i])
		// Change direction at first or last row
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	// Combine all rows
	result := []byte{}
	for _, row := range rows {
		result = append(result, row...)
	}
	return string(result)
}

func main() {
	s := "PAYPALISHIRING"
	fmt.Println(convert(s, 3)) // "PAHNAPLSIIGYIR"
}

/*
	| i  | s[i] | curRow | goingDown Before | Append To Row | goingDown After | curRow After |
	| -- | ---- | ------ | ---------------- | ------------- | --------------- | ------------ |
	| 0  | P    | 0      | false            | Row0          | true            | 1            |
	| 1  | A    | 1      | true             | Row1          | true            | 2            |
	| 2  | Y    | 2      | true             | Row2          | false           | 1            |
	| 3  | P    | 1      | false            | Row1          | false           | 0            |
	| 4  | A    | 0      | false            | Row0          | true            | 1            |
	| 5  | L    | 1      | true             | Row1          | true            | 2            |
	| 6  | I    | 2      | true             | Row2          | false           | 1            |
	| 7  | S    | 1      | false            | Row1          | false           | 0            |
	| 8  | H    | 0      | false            | Row0          | true            | 1            |
	| 9  | I    | 1      | true             | Row1          | true            | 2            |
	| 10 | R    | 2      | true             | Row2          | false           | 1            |
	| 11 | I    | 1      | false            | Row1          | false           | 0            |
	| 12 | N    | 0      | false            | Row0          | true            | 1            |
	| 13 | G    | 1      | true             | Row1          | true            | 2            |
*/
