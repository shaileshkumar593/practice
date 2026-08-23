package main

import (
	"fmt"
	"time"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	go func() {
		panic("boom")
	}()

	time.Sleep(time.Second)
}

/*

Because panic happens inside a different goroutine.

This is the key rule:

recover() can only recover a panic that occurs in the same goroutine where that recover() is executing.

Your code has two goroutines.
*/
