package main

import (
	"fmt"
	"sync"
	"time"
)

func worker1(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(time.Duration(<-jobs) * time.Second)
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	// Start 5 workers
	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go worker1(w, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= 100; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}

/*

	core concurrency scheduling model of Go.
	Let’s dissect exactly how a job from a channel gets assigned (“allocated”) to a worker.

	 Step-by-Step: How a Job Is Allocated to a Worker
1️⃣ Setup — Workers Waiting
After the 5 goroutines start:

go worker(1, jobs, &wg)
go worker(2, jobs, &wg)
...
Each of them immediately hits this line inside the function:

for job := range jobs {
Here’s what happens internally:

Each worker calls <-jobs to receive from the channel.

Since no job is available yet, all 5 goroutines are blocked on the same channel.

The Go runtime keeps them in a wait queue associated with jobs.

So we have:

Worker	State	Waiting on
Worker 1	Blocked	jobs channel
Worker 2	Blocked	jobs channel
Worker 3	Blocked	jobs channel
Worker 4	Blocked	jobs channel
Worker 5	Blocked	jobs channel
2️⃣ Job Sent by Main Goroutine
When main sends the first job:

jobs <- 1
Here’s what happens inside the runtime:

The channel jobs checks its receiver queue (workers waiting).

Finds one blocked worker (say Worker 1).

Transfers the job (1) directly from sender to receiver without storing it (unbuffered channel = direct handoff).

Worker 1 is unblocked and resumes running.

The main goroutine continues to the next send (jobs <- 2).

So the flow is:

main → jobs <- 1 → wake Worker 1 → Worker 1 handles job 1
3️⃣ Subsequent Jobs
Each time main sends another job:

jobs <- j
The channel checks if any worker is waiting.

It picks one waiting worker from the queue.

That worker receives the job and starts executing.

Once the worker finishes the iteration (after fmt.Printf),
it loops back to for job := range jobs → and waits for the next job again.

This creates a natural balancing mechanism:

Idle workers are always waiting.

The runtime gives each new job to a waiting worker.

4️⃣ If All Workers Are Busy
Now suppose all 5 workers are currently processing jobs (e.g., heavy tasks).

Then, when main executes:

jobs <- j
The channel has no waiting receivers.

Because the channel is unbuffered, this send blocks.

Main goroutine is paused until a worker finishes and comes back to receive.

Then:

A worker finishes a job.

Loops back and executes <-jobs again.

That worker is now ready.

The blocked sender (jobs <- j) and this ready receiver are paired.

Job is delivered immediately.

Both continue.

✅ No polling, no spin, no wasted CPU — pure synchronization.

5️⃣ Closing the Channel
When all 20 jobs are sent:

close(jobs)
Channel signals no more sends allowed.

Any worker doing:

for job := range jobs
will:

Keep consuming pending jobs.

Exit loop gracefully once the channel is empty and closed.

🧠 Visualization of Channel Matching
At runtime, this process looks like:

            ┌────────────────────────────┐
            │         jobs chan          │
            └────────────────────────────┘
       Receivers (workers) waiting here

┌───────────────┐
│ main goroutine│
│ sends jobs -> │
└───────────────┘
       ↓
(1) jobs <- 1 → wakes Worker 1
(2) jobs <- 2 → wakes Worker 2
(3) jobs <- 3 → wakes Worker 3
...
(6) if all busy → main blocks
(7) when a worker done → loops → receives new job
🧩 Scheduling Notes
Go’s scheduler runs on top of OS threads (M:N model):

M = OS threads

G = goroutines

P = logical processors (scheduling context)

When a worker blocks on <-jobs, it’s parked.

Another runnable goroutine takes its place — no thread blocking.

When a job arrives, the parked worker is resumed by the scheduler.

So the scheduler guarantees:

Each worker gets a job only when available.

Each job is processed by exactly one worker.

There is no contention or spin.

⚖️ Summary
Step	Action	Behavior
1	Start 5 workers	All block on <-jobs
2	Main sends a job	Wakes one waiting worker
3	Worker processes job	Then blocks again waiting for next
4	If all busy	Sender (main) blocks until one frees
5	Close channel	Workers exit gracefully after consuming all
🧩 Analogy
Imagine:

5 chefs (workers) standing idle in a kitchen.

Orders (jobs) come in one at a time.

As soon as an order arrives → the first available chef grabs it.

When all chefs are busy → new orders wait.

When orders stop coming (channel closed) → chefs finish remaining ones and go home.


	                  +--------------------+
                  |       main()       |
                  |--------------------|
                  | create 5 workers   |
                  | send 20 jobs       |
                  | close(jobs)        |
                  | wg.Wait()          |
                  +---------+----------+
                            |
                            ↓
       +------------------------------------------+
       |                 jobs chan                |
       +------------------------------------------+
              ↑        ↑        ↑        ↑        ↑
              |        |        |        |        |
     +---------+   +---------+  +---------+  +---------+  +---------+
     | Worker 1|   | Worker 2|  | Worker 3|  | Worker 4|  | Worker 5|
     +---------+   +---------+  +---------+  +---------+  +---------+
         ↓              ↓             ↓             ↓             ↓
  Process job 1    Process job 2  Process job 3  Process job 4  Process job 5
  Process job 6    Process job 7  ... etc until all 20 jobs done



*/
