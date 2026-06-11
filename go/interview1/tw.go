package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	mu         sync.Mutex
	timestamps []time.Time
	limit      int
	window     time.Duration
}

func NewLimiter(limit int, window time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		limit:  limit,
		window: window,
	}
}

func (s *SlidingWindowLimiter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-s.window)

	// Remove old timestamps
	idx := 0
	for _, t := range s.timestamps {
		if t.After(cutoff) {
			break
		}
		idx++
	}
	s.timestamps = s.timestamps[idx:]

	// Check limit
	if len(s.timestamps) < s.limit {
		s.timestamps = append(s.timestamps, now)
		return true
	}
	return false
}

func Printvalue(c chan int, wg *sync.WaitGroup, id int, limiter *SlidingWindowLimiter) {
	defer wg.Done()

	for v := range c {

		// 🔥 wait until allowed
		for !limiter.Allow() {
			time.Sleep(10 * time.Millisecond)
		}

		fmt.Println("Worker", id, ":", v)
	}
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	// Allow 2 requests per second
	limiter := NewLimiter(2, time.Second)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Printvalue(c, &wg, i, limiter)
	}

	// Producer
	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()

	wg.Wait()
}