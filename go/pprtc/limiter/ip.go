package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// store request timestamps per IP
type clientData struct {
	timestamps []time.Time
	count      int // 👈 added count
}

var (
	mu      sync.Mutex
	clients = make(map[string]*clientData)
	limit   = 5
	window  = time.Minute
)

// Middleware
func rateLimiter(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		now := time.Now()

		mu.Lock()
		defer mu.Unlock()

		if _, exists := clients[ip]; !exists {
			clients[ip] = &clientData{}
		}

		client := clients[ip]

		// 1️⃣ Remove old timestamps
		valid := []time.Time{}
		for _, t := range client.timestamps {
			if now.Sub(t) < window {
				valid = append(valid, t)
			}
		}
		client.timestamps = valid

		// 👇 Update count after cleanup
		client.count = len(client.timestamps)

		// 2️⃣ Check limit
		if client.count >= limit {
			http.Error(w,
				fmt.Sprintf("Too Many Requests. Count: %d", client.count),
				http.StatusTooManyRequests)
			return
		}

		// 3️⃣ Add current request
		client.timestamps = append(client.timestamps, now)

		// 👇 Increment count
		client.count++

		// 4️⃣ Add count in response header
		w.Header().Set("X-Request-Count", fmt.Sprintf("%d", client.count))

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