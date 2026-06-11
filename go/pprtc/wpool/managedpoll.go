package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	numWorkers := runtime.NumCPU() // scale with CPU cores
	jobs := make(chan int)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing %d\n", id, job)
			}
		}(w)
	}

	for j := 1; j <= 50; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}

/*
👉 Scales based on available CPUs.
👉 Ideal for high-performance concurrent workloads.*/
