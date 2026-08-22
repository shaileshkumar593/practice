package main

import "fmt"

func main() {
	var c chan int

	v := <-c
	fmt.Print(v)
}
