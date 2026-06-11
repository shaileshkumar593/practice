package main

import (
	"fmt"
	"sync"
)

//Like the name suggests, WaitGroup is a primitive that allows us to wait for a group of goroutines to finish.

func main() {
	wg := &sync.WaitGroup{}
	// We need to wait for one goroutine, so we Add(1) *before* starting it.
	wg.Add(1)
	go func() {
		// Defer Done() right away to ensure it's called, even if the goroutine panics or returns early.
		defer wg.Done()
		fmt.Println("Hello, World!")
	}()
	wg.Wait() // Wait blocks until the counter is zero.
}
