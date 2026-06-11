package main

import (
	"fmt"
	"time"
)

// fan-out problem

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)

		result := job * 2

		fmt.Printf("worker %d finished job %d\n", id, job)
		results <- result
	}
}

func main() {
	numjobs := 5

	jobs := make(chan int, numjobs)
	results := make(chan int, numjobs)

	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}

	for j := 1; j <= numjobs; j++ {
		jobs <- j
	}

	close(jobs)

	for i := 1; i <= numjobs; i++ {
		res := <-results
		fmt.Println("Result : ", res)
	}

}
