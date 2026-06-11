package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		results <- job * 2
		fmt.Printf("Worker %d processed %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	// fan-out workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// send jobs
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// close results after workers done
	go func() {
		wg.Wait()
		close(results)
	}()

	// fan-in results
	for result := range results {
		fmt.Println("Result:", result)
	}
}

/*

	🔹 Fan-Out
🧠 Idea
		One producer distributes work to multiple workers.

				---> Worker1
		Producer ---> Worker2
				---> Worker3
🔹 Fan-In
🧠 Idea
		Multiple workers send results back into one channel.

		Worker1 ---
		Worker2 ----> Results Channel
		Worker3 ---


*/
