package main

import (
	"fmt"
	"time"

	"github.com/zenthangplus/goccm"
)

func main() {
	// Limit 3 goroutines to run concurrently.
	c := goccm.New(3)

	for i := 1; i <= 10; i++ {

		// This function has to call before any goroutine
		c.Wait()

		go func(i int) {
			fmt.Printf("Job %d is running\n", i)
			time.Sleep(2 * time.Second)

			// This function has to when a goroutine has finished
			// Or you can use `defer c.Done()` at the top of goroutine.
			c.Done()
		}(i)
	}

	// This function has to call to ensure all goroutines have finished
	// after close the main program.
	c.WaitAllDone()
}
