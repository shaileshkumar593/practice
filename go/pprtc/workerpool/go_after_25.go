package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		// The Go method adds 1 to the counter and calls Done() on completion.
		wg.Go(func() {
			worker(i)
		})
	}
	// Block until all goroutines launched by wg.Go() are done.
	wg.Wait()

	fmt.Println("All workers finished")
}
