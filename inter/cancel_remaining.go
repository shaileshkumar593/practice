package main

import (
	"context"
	"fmt"
	"time"
)

// First Result Wins (Cancel Remaining Goroutines)

func worker(
	ctx context.Context,
	id int,
	result chan<- string,
) {

	select {

	case <-time.After(time.Duration(id) * time.Second):
		result <- fmt.Sprintf("worker %d won", id)

	case <-ctx.Done():
		return
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	result := make(chan string)

	for i := 1; i <= 5; i++ {
		go worker(ctx, i, result)
	}

	winner := <-result

	cancel()

	fmt.Println(winner)
}
