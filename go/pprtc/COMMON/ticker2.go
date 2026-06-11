package main

import (
	"fmt"
	"time"
)

func main() {
	queue := []string{"Job1", "Job2", "Job3"}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if len(queue) == 0 {
				fmt.Println("No jobs left — stopping polling.")
				return
			}
			job := queue[0]
			queue = queue[1:]
			fmt.Println("Processing:", job, "at", time.Now().Format("15:04:05"))
		}
	}
}

/*

	| Field / Function                  | Description                                                      |
| --------------------------------- | ---------------------------------------------------------------- |
| `time.NewTicker(d time.Duration)` | Creates a ticker that sends the current time every `d` duration. |
| `ticker.C`                        | A channel where time values are delivered periodically.          |
| `ticker.Stop()`                   | Stops the ticker and releases associated resources.              |
| `range ticker.C`                  | Continuously reads from the channel until stopped.               |

*/
