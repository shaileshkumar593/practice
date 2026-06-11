package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Starting goroutines...")

	wg.Go(func() { // Using the new WaitGroup.Go() method
		fmt.Println("Goroutine 1 running...")
		time.Sleep(100 * time.Millisecond)
	})

	wg.Go(func() { // Using the new WaitGroup.Go() method
		fmt.Println("Goroutine 2 running...")
		time.Sleep(150 * time.Millisecond)
	})

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("All goroutines finished.")
}
