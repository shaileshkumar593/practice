package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	ch <- 100

	close(ch)

	fmt.Println(<-ch)

	fmt.Println(<-ch) // Second receive returns Zero value.
}