package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

/*
SlidingWindowLimiter implements
Strict "N requests in last T duration"
*/
type SlidingWindowLimiter struct {
	mu        sync.Mutex
	requests  []time.Time
	limit     int
	windowDur time.Duration
}

// Constructor
func NewSlidingWindowLimiter(limit int, window time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		limit:     limit,
		windowDur: window,
		requests:  make([]time.Time, 0, limit),
	}
}

/*
Allow checks if request is allowed.

Algorithm:
1. Get current time
2. Remove expired timestamps (older than window)
3. If count >= limit → reject
4. Else append and allow
*/
func (s *SlidingWindowLimiter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-s.windowDur)

	// Remove expired timestamps from front
	idx := 0
	for idx < len(s.requests) && s.requests[idx].Before(cutoff) {
		idx++
	}

	// Shrink slice (O(1) re-slice)
	s.requests = s.requests[idx:]

	// Check limit
	if len(s.requests) >= s.limit {
		return false
	}

	// Add current request
	s.requests = append(s.requests, now)
	return true
}

func main() {
	// 10 requests in last 60 seconds
	limiter := NewSlidingWindowLimiter(10, 60*time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		fmt.Fprintf(w, "Request allowed at %v\n", time.Now())
	})

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*

	Let’s break this struct down field by field and understand why each exists.

type SlidingWindowLimiter struct {
	mu        sync.Mutex
	requests  []time.Time
	limit     int
	windowDur time.Duration
}
This struct represents a strict sliding window rate limiter.

It enforces:

“At most limit requests in the last windowDur duration.”

🔐 1️⃣ mu sync.Mutex
What it is:
A mutual exclusion lock.

Why we need it:
Multiple goroutines may call Allow() at the same time.

Without mutex:

Two goroutines may modify requests slice together

Slice corruption

Data race

Panic

Incorrect rate limiting

What it protects:
It protects:

requests []time.Time
Because slice operations are not thread-safe.

🕒 2️⃣ requests []time.Time
What it stores:
All request timestamps that happened within the window.

Example:

If window = 60 seconds and limit = 3

At time = 100s

requests = [70s, 85s, 95s]
These are the timestamps of previous requests.

Why we store timestamps:
To implement strict rolling window.

Every new request:

Remove timestamps older than now - windowDur

Check if remaining count >= limit

If yes → reject

Else → append current timestamp

Why slice is efficient here?
Because:

Requests come in chronological order

Oldest request always at index 0

We only remove from front

Append at end

So it behaves like a queue.

🎯 3️⃣ limit int
What it represents:
Maximum number of requests allowed inside the window.

Example:

limit = 10
Means:

Maximum 10 requests in last windowDur.

⏳ 4️⃣ windowDur time.Duration
What it represents:
Time window size.

Example:

windowDur = 60 * time.Second
Means:

Only requests from last 60 seconds count.

🧠 How All Fields Work Together
Suppose:

limit = 3
windowDur = 10 seconds
Timeline:

Time	Action	requests slice
0s	allow	[0]
3s	allow	[0,3]
5s	allow	[0,3,5]
6s	reject	[0,3,5]
11s	cleanup 0	[3,5]
11s	allow	[3,5,11]
The struct stores only what is needed.

🔥 Important Design Properties
1️⃣ Strict Rolling Window
Unlike token bucket:

This guarantees:

At ANY moment, there are never more than limit requests in the last windowDur.

Exact behavior.

2️⃣ Memory Bound
Max size of:

requests <= limit
So memory usage is predictable.

3️⃣ Thread Safe
Because of:

mu sync.Mutex
Without it → unsafe.

⚠️ What Happens If We Remove Mutex?
If 1000 goroutines call Allow():

Concurrent append

Concurrent slice reslicing

Runtime panic

Race condition

You’ll see:

fatal error: concurrent map/slice writes
🏗️ Conceptual Model
Think of this struct as:

Queue of timestamps + Lock + Rules
The lock ensures:

Only one request checks/modifies the queue at a time.

📦 Why Not Use Channel Instead?
Because:

We need to remove old timestamps

We need random access

Channel does not allow inspection/removal of older elements

So slice is simpler.

🚀 Advanced Improvement (For High Traffic)
Instead of:

[]time.Time
You can use:

Ring buffer

Circular queue

Atomic index pointer

Redis sorted set (distributed)

🏁 Final Summary
This struct holds:

Field	Purpose
mu	Protect shared state
requests	Store timestamps inside window
limit	Max allowed requests
windowDur	Time window size
Together, they implement:

Strict N requests in last T seconds rate limiter.

*/
