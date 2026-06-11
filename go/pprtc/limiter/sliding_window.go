package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	mu        sync.Mutex
	requests  []time.Time
	limit     int
	windowDur time.Duration
}

func NewSlidingWindowLimiter(limit int, window time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		limit:     limit,
		windowDur: window,
		requests:  make([]time.Time, 0),
	}
}

func (s *SlidingWindowLimiter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-s.windowDur)
	fmt.Println("--------------------- ", cutoff)

	// Remove expired timestamps
	valid := make([]time.Time, 0)
	for _, t := range s.requests {
		if t.After(cutoff) {
			valid = append(valid, t)
		}
	}
	s.requests = valid

	// Check limit
	if len(s.requests) >= s.limit {
		return false
	}

	// Add current request
	s.requests = append(s.requests, now)
	return true
}

func main() {
	limiter := NewSlidingWindowLimiter(10, 60*time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		fmt.Fprintf(w, "Request allowed at %v\n", time.Now())
	})

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
	1. curtime :=  time.now
	2. expired := curtime.Add(-s.windowDur)


*/
/*
In Rate Limiter Context
You often see:

cutoff := time.Now().Add(-60 * time.Second)

if requestTime.After(cutoff) {
    // keep request (within last 60 sec)
}
This means:

Only keep requests that happened AFTER (later than) the cutoff time.
*/
