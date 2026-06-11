/*
	If interviewer asks “write 3 producer 3 consumer in Go”, write:

one shared jobs channel

3 producer goroutines

3 consumer goroutines

producer waitgroup

consumer waitgroup

close channel after all producers are done
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(id int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		job := id*100 + i
		fmt.Printf("Producer %d produced job %d\n", id, job)
		jobs <- job
		time.Sleep(300 * time.Millisecond)
	}
}

func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Consumer %d consumed job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("Consumer %d exiting\n", id)
}

func main() {
	jobs := make(chan int, 5)

	var producerWG sync.WaitGroup
	var consumerWG sync.WaitGroup

	// Start 3 consumers
	for i := 1; i <= 3; i++ {
		consumerWG.Add(1)
		go consumer(i, jobs, &consumerWG)
	}

	// Start 3 producers
	for i := 1; i <= 3; i++ {
		producerWG.Add(1)
		go producer(i, jobs, &producerWG)
	}

	// Wait for all producers to finish
	producerWG.Wait()

	// Close channel so consumers stop
	close(jobs)

	// Wait for all consumers to finish
	consumerWG.Wait()

	fmt.Println("All producers and consumers completed")
}