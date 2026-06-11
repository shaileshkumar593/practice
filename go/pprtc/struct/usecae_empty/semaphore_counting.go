package main

import (
	"fmt"
	"time"
)

func main() {
	sem := make(chan struct{}, 2)

	for i := 1; i <= 5; i++ {
		sem <- struct{}{} // acquire

		go func(id int) {
			fmt.Println("Running", id)
			time.Sleep(2 * time.Second)
			<-sem // release
		}(i)
	}

	time.Sleep(6 * time.Second)
}

/*

Execution Flow (VERY IMPORTANT)
We run loop from 1 → 5.

🟢 Iteration 1 (i=1)
sem <- struct{}{}   // buffer size becomes 1
Goroutine 1 starts.

Buffer:

[ x ]   (1/2 full)
🟢 Iteration 2 (i=2)
sem <- struct{}{}   // buffer size becomes 2
Goroutine 2 starts.

Buffer:

[ x x ]   (2/2 full)
🔴 Iteration 3 (i=3)
sem <- struct{}{}
⚠️ Buffer is FULL.

Channel capacity = 2
Currently = 2

So this line BLOCKS.

Main goroutine waits here.

🚨 No third goroutine starts yet.

⏳ After 2 Seconds
Goroutine 1 or 2 finishes:

<-sem   // release
One slot becomes free.

Buffer:

[ x ]   (1/2 full)
Now main goroutine unblocks.

🟢 Iteration 3 resumes
Now:

sem <- struct{}{}
Goroutine 3 starts.

Buffer:

[ x x ]   (2/2 full)
Same process continues for 4 and 5.

📊 Concurrency Visualization
Time 0s:

Running 1
Running 2
Time 2s:

1 finishes
2 finishes
Running 3
Running 4
Time 4s:

3 finishes
4 finishes
Running 5
Time 6s:

5 finishes
🎯 What Is Happening Conceptually?
You are implementing:

Max 2 concurrent workers
This is called:

Counting Semaphore

🧠 Why Does sem <- struct{}{} Block?
Because:

Sending to a buffered channel blocks if buffer is full

Receiving blocks if buffer is empty

That blocking behavior is used intentionally.

⚠ Why time.Sleep(6 * time.Second) in main?
Because:

main() exits immediately after loop

If main exits → program exits → goroutines die

So we wait 6 seconds to allow all goroutines to finish.

🚨 Production-Grade Version (Better)
Using sync.WaitGroup instead of sleep:

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sem := make(chan struct{}, 2)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		sem <- struct{}{}
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			defer func() { <-sem }()

			fmt.Println("Running", id)
			time.Sleep(2 * time.Second)
		}(i)
	}

	wg.Wait()
}
This is the correct production pattern.

🔥 Why This Pattern Is Powerful
Used in:

API rate limiting

Worker pools

Database connection limiting

Limiting file processing

Controlling CPU-intensive tasks

🧠 Mental Model
Think of channel buffer like:

Parking slots = 2
Cars (goroutines) can park only if slot free
When slot full → next car waits.

📊 Complexity
Time complexity: depends on work
Space complexity: O(capacity)

🏆 Interview-Level Explanation
If interviewer asks:

"How do you limit concurrency in Go without mutex?"

Answer:

"Using buffered channel as semaphore."

Explain:

Channel capacity = max concurrency

Send = acquire

Receive = release

struct{} used for zero memory

🔥 Advanced Insight
This is how Go internally implements:

Worker pools

Concurrency throttling

Backpressure systems

*/
