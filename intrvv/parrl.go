package main

import (
	"fmt"
	"runtime"
	"sync"
)

func calculate(n int) uint64 {
	var result uint64

	for i := 0; i < 100_000_000; i++ {
		result += uint64((i * n) % 100)
	}

	return result
}

func worker(
	id int,
	jobs <-chan int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for job := range jobs {
		result := calculate(job)

		fmt.Printf(
			"Worker %d completed job %d => %d\n",
			id,
			job,
			result,
		)
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	jobs := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)

		go worker(i, jobs, &wg)
	}

	for i := 1; i <= 8; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	fmt.Println("All jobs completed")
}
