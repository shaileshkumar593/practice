package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a pool
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Allocating new object")
			return make([]byte, 8)
		},
	}

	// Get object from pool
	obj1 := pool.Get().([]byte)
	fmt.Println("Using obj1")
	pool.Put(obj1) // Return to pool

	// Get again
	obj2 := pool.Get().([]byte)
	fmt.Println("Using obj2")
	pool.Put(obj2)
}
