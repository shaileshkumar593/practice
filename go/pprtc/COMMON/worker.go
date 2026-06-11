package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		default:
			value := ctx.Value("key")
			fmt.Println(value)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	fmt.Println("Go Context Tutorial")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for w := 1; w <= 3; w++ {
		go worker(ctx, w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
	select {
	case <-ctx.Done():
		fmt.Println("Took too long!")
	default:
	}
	time.Sleep(5 * time.Second)

}
