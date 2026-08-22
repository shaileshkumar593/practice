package main

import "fmt"

func main() {
	x := 10

	defer func() {
		fmt.Println(x)
	}()

	x = 20
}

/*
	return expression evaluated
        ↓
named return variables updated
        ↓
defer functions execute
        ↓
function returns to caller
*/
