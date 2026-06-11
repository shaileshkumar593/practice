package main

import (
	"fmt"
	"time"
)

func performTasks(ch chan string) {
	// Simulate a task that takes 2 seconds
	time.Sleep(2 * time.Second)
	ch <- "Task completed"
}

func main() {
	taskChannel := make(chan string)
	go performTasks(taskChannel)

	for i := 0; i <= 5; i++ {
		select {
		case msg := <-taskChannel:
			fmt.Println("Received:", msg)
		case <-time.After(1 * time.Second): // Timeout after 1 second
			fmt.Println("Timeout: Task took too long")
		}
	}
}

/*
1. Goroutine Starts: The performTasks function is launched in a separate goroutine. It immediately starts a 2-second sleep.
The Loop: The main function enters a for loop that iterates 5 times.
2. First Iteration (at ~0s): The select statement waits for 1 second (time.After(1 * time.Second)). Since the task takes 2 seconds, the timeout case executes first.
Output: Timeout: Task took too long
3. Second Iteration (at ~1s): The select statement waits again. The task is still running (it finishes at the 2-second mark). The timeout case executes again.
Output: Timeout: Task took too long
4. Task Completion (at ~2s): The performTasks goroutine finishes its sleep and sends the message "Task completed" to taskChannel.
5. Third Iteration (at ~2s): The select statement is executed again. The taskChannel now has a message waiting. The case msg := <-taskChannel: branch is selected immediately.
Output: Received: Task completed
6. Fourth and Fifth Iterations (at ~3s and ~4s): The loop continues. The taskChannel is now empty. The select statements in
 these iterations will timeout after 1 second, executing the timeout case.
Output: Timeout: Task took too long (twice more)
---------  The exact ordering of the first two "Timeout" messages relative to the internal scheduling might vary slightly
depending on system load, but the pattern of two timeouts, one receipt, and two more timeouts is the intended behavior.
*/
