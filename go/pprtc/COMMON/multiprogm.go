package main

import "github.com/zenthangplus/goccm"

func main() {
	// Create the concurrency manager
	// The first argument is the maximum number of goroutines to run concurrently.
	c := goccm.New(10)

	// Wait until a slot is available for the new goroutine.
	c.Wait()

	// Mark a goroutine as finished
	c.Done()

	// Wait for all goroutines are done
	c.WaitAllDone()

	// Close the manager manually
	c.Close()

	// Returns the number of goroutines which are running
	c.RunningCount()
}
