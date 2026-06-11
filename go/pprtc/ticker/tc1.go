| Use Case                                 | Why Ticker?                                |
| ---------------------------------------- | ------------------------------------------ |
| Heartbeats / keep-alive ping             | Regular, repeating interval                |
| Periodic tasks (cleanup, refresh, sync)  | Executes repeatedly with a stable schedule |
| Monitoring and metrics                   | Emit metrics every N seconds               |
| Rate-limited background jobs             | Execute small batch every tick             |
| Polling external APIs                    | Check status at fixed intervals            |
| Time-based worker control                | Workers wake up every interval             |
| Graceful shutdown of periodic goroutines | Combine ticker with context                |


| Property   | Description                                    |
| ---------- | ---------------------------------------------- |
| Interval   | Duration between ticks (`time.Duration`)       |
| Channel    | `ticker.C` channel sends `time.Time` values    |
| Stoppable  | `ticker.Stop()` stops the ticker               |
| Resettable | `ticker.Reset(d)` changes interval dynamically |



package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go heartbeat(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	// Give goroutine time to exit
	time.Sleep(200 * time.Millisecond)
}

func heartbeat(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() // prevents memory leaks

	for {
		select {
		case <-ticker.C:
			fmt.Println("❤️ heartbeat sent")
		case <-ctx.Done():
			fmt.Println("heartbeat stopped:", ctx.Err())
			return
		}
	}
}
/*
	A service sends a heartbeat signal (e.g., to another server, a queue, or logs).
	The heartbeat should stop cleanly when the context is canceled.

*/
