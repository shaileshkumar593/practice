package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++ // ❌ unsafe
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final:", counter)
}

/*
	 Solution 1: Use Mutex (Safe)
var mu sync.Mutex

go func() {
	mu.Lock()
	counter++
	mu.Unlock()
}()
✔ Only one goroutine modifies at a time
✔ Correct result
❌ Slightly slower (locking cost)

✅ Solution 2: Use Atomic (Best for Counter)
import "sync/atomic"

var counter int64

go func() {
	atomic.AddInt64(&counter, 1)
}()
✔ No mutex
✔ Lock-free
✔ Very fast
✔ Production-ready

⚡ Why Atomic Is Better for Counters
Because CPU provides hardware-level atomic instructions
*/
