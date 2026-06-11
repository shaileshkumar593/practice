package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		id := i
		wg.Go(func() { // Using WaitGroup.Go()
			worker(id)
		})
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
