package main

import (
	"fmt"
	"sync"
	"time"
)

// ---------- Producer ----------
func producer(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			fmt.Println("produced:", n)
			out <- n
			time.Sleep(300 * time.Millisecond)
		}
	}()

	return out
}

// ---------- Worker (Fan-Out) ----------
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d processing %d\n", id, job)

		time.Sleep(500 * time.Millisecond)

		results <- job * 2
	}
}

// ---------- Fan-Out ----------
func fanOut(jobs <-chan int, workerCount int) <-chan int {
	results := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Close results after all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// ---------- Main ----------
func main() {
	jobs := producer([]int{1, 2, 3, 4, 5})

	results := fanOut(jobs, 3)

	for r := range results {
		fmt.Println("result:", r)
	}

	fmt.Println("done")
}
