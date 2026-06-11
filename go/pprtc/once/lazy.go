package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// SomeType represents an expensive resource (e.g., DB connection, config, cache)
type SomeType struct {
	Timestamp time.Time
	Message   string
}

// Singleton variables
var (
	once              sync.Once
	expensiveResource *SomeType
)

// initializeExpensiveResource simulates a costly initialization (e.g., setup, file IO, DB conn)
func initializeExpensiveResource() *SomeType {
	log.Println("Initializing expensive resource...")
	time.Sleep(2 * time.Second) // Simulate delay
	return &SomeType{
		Timestamp: time.Now(),
		Message:   "Singleton initialized",
	}
}

// getExpensiveResource uses sync.Once to initialize resource only once
func getExpensiveResource() *SomeType {
	once.Do(func() {
		expensiveResource = initializeExpensiveResource()
	})
	return expensiveResource
}

func main() {
	var wg sync.WaitGroup

	// Simulate multiple goroutines accessing the resource
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			res := getExpensiveResource()
			log.Printf("[Goroutine %d] Resource: %v\n", id, res)
		}(i)
	}

	wg.Wait()

	fmt.Println("All goroutines finished.")
}


/*
Key Features:
✅ sync.Once ensures thread-safe lazy initialization

✅ Only one goroutine initializes the resource

✅ Useful for expensive/shared resource loading

✅ Clean exit via sync.WaitGroup

✅ No race conditions

Why Not Use init()?
Using init() or global assignment:

Would initialize the resource eagerly, during program startup.

Even if the resource is never used, the cost is still paid.

With lazy initialization:

Initialization is deferred until first real use.
Thread-Safe + Lazy
This pattern is essential when:

Initialization is expensive (DB connection, file IO, configuration)

You need concurrent-safe singleton in Go

You're avoiding redundant computation or resource contention







*/