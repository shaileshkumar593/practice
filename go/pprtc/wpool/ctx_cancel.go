package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // 🔴 1️⃣ Check for cancellation
			fmt.Printf("Worker %d exiting\n", id)
			return

		case job, ok := <-jobs: // 🟢 2️⃣ Get job if available
			if !ok {
				return // channel closed — end worker
			}
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(time.Second) // simulate work
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // 🔹 1️⃣ Create cancellable context
	jobs := make(chan int)
	var wg sync.WaitGroup

	// 🔹 2️⃣ Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, &wg)
	}

	// 🔹 3️⃣ Send 10 jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	// 🔹 4️⃣ Wait for a bit, then cancel context (simulate shutdown)
	time.Sleep(3 * time.Second)
	cancel() // ❗Triggers graceful exit for all workers

	close(jobs)
	wg.Wait()
	fmt.Println("✅ All workers stopped gracefully.")
}

/*
	 Step-by-Step Explanation
1️⃣ Context Creation
ctx, cancel := context.WithCancel(context.Background())
context.Background() is the root context.

context.WithCancel returns:

A derived context ctx

A cancel function to broadcast cancellation

When you call cancel(), every goroutine using this context receives a signal via ctx.Done() channel.

2️⃣ Worker Start
Each worker starts and enters an infinite loop:

for {
	select {
	case <-ctx.Done():
		// shutdown signal received
	case job, ok := <-jobs:
		// process jobs
	}
}
🧩 Internally:

The select statement waits for whichever case becomes ready.

Either:

A job arrives on jobs, or

The context gets cancelled (ctx.Done() becomes readable)

Whichever happens first, that branch executes.

3️⃣ Job Processing
case job, ok := <-jobs:
	if !ok {
		return // channel closed
	}
	fmt.Printf("Worker %d processing job %d\n", id, job)
	time.Sleep(time.Second)
If a job is received:

Worker processes it (simulated by time.Sleep).

Loops again to check for the next job or cancellation.

4️⃣ Cancellation Signal
After 3 seconds:

cancel()
cancel() closes the ctx.Done() channel.

Every worker that’s blocked in the select immediately unblocks.

The case:

case <-ctx.Done():
becomes active, triggering:

fmt.Printf("Worker %d exiting\n", id)
return
Workers exit gracefully — no panic, no abrupt termination.

5️⃣ Channel Close and WaitGroup
close(jobs)
wg.Wait()
fmt.Println("✅ All workers stopped gracefully.")
close(jobs) prevents further sends.

wg.Wait() waits for all workers to call wg.Done() (inside defer).

Ensures no goroutine leak.

Program exits cleanly only when all workers are finished.

Where This Pattern Is Used in Production
HTTP servers → cancel all request goroutines on shutdown.

Pipelines / ETL jobs → stop long-running data ingestion cleanly.

Microservices → cancel child tasks on parent termination.

Worker pools → graceful draining on signal interrupt (SIGINT/SIGTERM).

*/
