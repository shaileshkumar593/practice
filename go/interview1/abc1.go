package main

import (
	"fmt"
	"sync"
)

// 🔹 Producer: generates jobs
func producer(id int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		job := id*10 + i
		jobs <- job
		fmt.Println("Produced:", job)
	}
}

// 🔹 Consumer: processes jobs
func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Consumer %d processed %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 5) // buffer helps reduce blocking

	var prodWG sync.WaitGroup
	var consWG sync.WaitGroup

	// 🔹 Start multiple producers
	for i := 1; i <= 3; i++ {
		prodWG.Add(1)
		go producer(i, jobs, &prodWG)
	}

	// 🔹 Start multiple consumers
	for i := 1; i <= 2; i++ {
		consWG.Add(1)
		go consumer(i, jobs, &consWG)
	}

	// 🔹 Close jobs AFTER all producers finish
	go func() {
		prodWG.Wait()
		close(jobs)
	}()

	// 🔹 Wait for all consumers to finish
	consWG.Wait()
}