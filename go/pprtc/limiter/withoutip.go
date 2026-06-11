package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// global request tracking
type requestData struct {
	timestamps []time.Time
	count      int
}

var (
	mu      sync.Mutex
	data    = &requestData{}
	limit   = 5
	window  = time.Minute
)

// Middleware
func rateLimiter(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		now := time.Now()

		mu.Lock()
		defer mu.Unlock()

		// 1️⃣ Remove old timestamps
		valid := []time.Time{}
		for _, t := range data.timestamps {
			if now.Sub(t) < window {
				valid = append(valid, t)
			}
		}
		data.timestamps = valid

		// 2️⃣ Update count
		data.count = len(data.timestamps)

		// 3️⃣ Check limit
		if data.count >= limit {
			http.Error(w,
				fmt.Sprintf("Too Many Requests (Global). Count: %d", data.count),
				http.StatusTooManyRequests)
			return
		}

		// 4️⃣ Add current request
		data.timestamps = append(data.timestamps, now)
		data.count++

		// 5️⃣ Add header
		w.Header().Set("X-Request-Count", fmt.Sprintf("%d", data.count))

		next.ServeHTTP(w, r)
	})
}

// Handler
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/ping", rateLimiter(http.HandlerFunc(pingHandler)))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}