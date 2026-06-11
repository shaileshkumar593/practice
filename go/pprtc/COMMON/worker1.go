package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const shortDuration = 3 * time.Second

func worker3(ctx context.Context, id int, jobs <-chan int) {
	fmt.Println("jobs :", jobs)
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		fmt.Println(j * rand.Intn(100))

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
		time.Sleep(2 * time.Second)
	}

}

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	const numJobs = 5
	jobs := make(chan int, numJobs)
	fmt.Println("Go Context Tutorial")

	for w := 1; w <= 3; w++ {
		go worker3(ctx, w, jobs)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	select {
	case <-ctx.Done():
		fmt.Println("Took too long!")
		fmt.Println("Err ", ctx.Err())

	default:
		fmt.Println("continue")
	}
	time.Sleep(2 * time.Second)
}
