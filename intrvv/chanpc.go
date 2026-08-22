package main

import "fmt"

func producer() <-chan int {

	c := make(chan int)
	var n int

	fmt.Println("enter generating limit")
	fmt.Scanf("%d", &n)
	go func() {
		defer close(c)
		for i := 0; i < 3*n; i = i + 2 {
			c <- i
		}
	}()

	return c
}

func main() {

	rslt := producer()

	for val := range rslt {
		fmt.Println(val)
	}
}
