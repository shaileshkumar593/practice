package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindow struct {
	timestamps []time.Time
	limit      int
	window     time.Duration
	mu         sync.Mutex
}

func NewSlidingWindow(limit int, window time.Duration) *SlidingWindow {
	return &SlidingWindow{
		limit:  limit,
		window: window,
	}
}

func (sw *SlidingWindow) allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-sw.window)

	// remove old timestamps
	idx := 0
	for _, t := range sw.timestamps {
		if t.After(cutoff) {
			break
		}
		idx++
	}
	sw.timestamps = sw.timestamps[idx:]

	if len(sw.timestamps) < sw.limit {
		sw.timestamps = append(sw.timestamps, now)
		return true
	}

	return false
}

// ======================
// Per User Sliding Window
// ======================

type SlidingLimiter struct {
	users    map[string]*SlidingWindow
	lastSeen map[string]time.Time
	mu       sync.Mutex
}

func NewSlidingLimiter() *SlidingLimiter {
	sl := &SlidingLimiter{
		users:    make(map[string]*SlidingWindow),
		lastSeen: make(map[string]time.Time),
	}

	go sl.cleanup()
	return sl
}

func (sl *SlidingLimiter) get(user string) *SlidingWindow {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	if sw, ok := sl.users[user]; ok {
		sl.lastSeen[user] = time.Now()
		return sw
	}

	sw := NewSlidingWindow(5, time.Second) // 5 req/sec
	sl.users[user] = sw
	sl.lastSeen[user] = time.Now()
	return sw
}

func (sl *SlidingLimiter) Allow(user string) bool {
	sw := sl.get(user)
	return sw.allow()
}

// cleanup inactive users
func (sl *SlidingLimiter) cleanup() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		sl.mu.Lock()
		for user, t := range sl.lastSeen {
			if time.Since(t) > 1*time.Minute {
				delete(sl.users, user)
				delete(sl.lastSeen, user)
			}
		}
		sl.mu.Unlock()
	}
}

// ======================
// Demo
// ======================

func main() {
	sl := NewSlidingLimiter()

	user := "user1"

	for i := 0; i < 10; i++ {
		if sl.Allow(user) {
			fmt.Println("Allowed", i)
		} else {
			fmt.Println("Blocked", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
