package main

import (
	"fmt"
	"sync"
)

func producer(id int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		job := id*10 + i
		jobs <- job
		fmt.Println("Produced:", job)
	}
}

func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Consumer", id, "processed", job)
	}
}

func main() {
	jobs := make(chan int, 5)

	var producerWG sync.WaitGroup
	var consumerWG sync.WaitGroup

	for i := 1; i <= 3; i++ {
		producerWG.Add(1)
		go producer(i, jobs, &producerWG)
	}

	for i := 1; i <= 2; i++ {
		consumerWG.Add(1)
		go consumer(i, jobs, &consumerWG)
	}

	producerWG.Wait()
	close(jobs)

	consumerWG.Wait()
}