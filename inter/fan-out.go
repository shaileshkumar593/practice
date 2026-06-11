package main

import (
	"fmt"
	"sync"
)

// One input channel → Multiple workers

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processed %d\n", id, job)
	}
}

func main() {

	jobs := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
}
