package main

/*
	Alternative: Using the new iter package (Go 1.23+)
For synchronous, single-threaded iteration, Go 1.23 introduced the iter package, which provides a more efficient, function-based iteration pattern. This avoids the overhead of goroutines and channels when true concurrency isn't needed.
Example: The new iter package*/

import (
	"fmt"
	"iter" // New package for standardized iterators
)

// MyNumberGenerator defines a function that can be used with a range loop.
func MyNumberGenerator(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i <= end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	// The range loop now directly works with the generator function.
	for num := range MyNumberGenerator(1, 5) {
		fmt.Printf("Received: %d\n", num)
	}
}

/*Use code with caution.

Output:
sh
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5*/
