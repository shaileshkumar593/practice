// example-7-pipelines/main.go
package main

import (
	"context"
	"fmt"
	"sync"
)

func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func worker342(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			// simulate work
			select {
			case out <- n * n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func merge(ctx context.Context, chans ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			select {
			case out <- v:
			case <-ctx.Done():
				return
			}
		}
	}
	wg.Add(len(chans))
	for _, c := range chans {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c1 := gen(ctx, 2, 3, 4, 5)
	// fan-out: start 3 workers reading from same channel
	w1 := worker342(ctx, c1)
	// fan-in: merge results (single worker for simplicity)
	out := merge(ctx, w1)
	for v := range out {
		fmt.Println("result:", v)
		// cancel early example
		if v > 15 {
			fmt.Println("main: canceling")
			cancel()
		}
	}
}
