package main

import "fmt"

func f() (x int) {
	defer func() {
		x++
	}()

	return 10
}

func main() {
	fmt.Println(f())
}

//  output :10  The return value is evaluated before deferred functions run.
