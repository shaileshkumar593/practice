package main

import (
	"fmt"
	"sync"
)

// Worker Pool (100 Jobs, 10 Workers)
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d -> Job %d\n", id, job)
	}
}

func main() {

	jobs := make(chan int, 100)

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 1; i <= 100; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
}
