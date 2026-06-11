package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	fmt.Println("----- BAD: time.After in loop (DON'T DO THIS) -----")
	go badExample()

	time.Sleep(2 * time.Second)

	fmt.Println("\n----- GOOD: time.NewTimer (SAFE) -----")
	goodTimerExample()

	fmt.Println("\n----- BEST: context.WithTimeout (API STYLE) -----")
	contextExample()

	fmt.Println("\n----- WORKER WITH TIMEOUT -----")
	workerWithTimeout()
}

//////////////////////////////////////////////////////////////
// ❌ BAD EXAMPLE (memory leak in loops)
//////////////////////////////////////////////////////////////

func badExample() {
	for i := 0; i < 3; i++ {
		select {
		case <-time.After(500 * time.Millisecond): // ❌ new timer each iteration
			fmt.Println("bad timeout", i)
		}
	}
}

//////////////////////////////////////////////////////////////
// ✅ GOOD EXAMPLE (time.NewTimer)
//////////////////////////////////////////////////////////////

func goodTimerExample() {
	timer := time.NewTimer(500 * time.Millisecond)
	defer timer.Stop()

	for i := 0; i < 3; i++ {

		timer.Reset(500 * time.Millisecond)

		select {
		case <-timer.C:
			fmt.Println("good timeout", i)
		}
	}
}

//////////////////////////////////////////////////////////////
// ✅ BEST EXAMPLE (context.WithTimeout)
//////////////////////////////////////////////////////////////

func contextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	select {
	case <-time.After(2 * time.Second): // simulate long work
		fmt.Println("work done")
	case <-ctx.Done():
		fmt.Println("context timeout:", ctx.Err())
	}
}

//////////////////////////////////////////////////////////////
// 🚀 WORKER WITH TIMEOUT (REAL USE CASE)
//////////////////////////////////////////////////////////////

func workerWithTimeout() {
	ch := make(chan string)

	// worker
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "work finished"
	}()

	// timeout control
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	select {
	case res := <-ch:
		fmt.Println("result:", res)

	case <-ctx.Done():
		fmt.Println("worker timeout:", ctx.Err())
	}
}
