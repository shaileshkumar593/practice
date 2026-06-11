// Implement a simple REST API server in Go with one endpoint /ping.
// Add a rate limiter middleware that allows only 5 requests per minute
// per client IP. If the limit is exceeded, return HTTP 429 (Too Many Requests).

/*func rateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr

		// TODO: Complete this func
		next.ServeHTTP(w, r)

	})
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/ping", rateLimiter(http.HandlerFunc(pingHandler)))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
*/

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// Sliding window limiter per IP
type client struct {
	timestamps []time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*client)
	limit   = 5
	window  = time.Minute
)

// Extract only IP (remove port)
func getIP(addr string) string {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return addr
	}
	return host
}

func rateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := getIP(r.RemoteAddr)
		now := time.Now()

		mu.Lock()
		c, exists := clients[ip]
		if !exists {
			c = &client{}
			clients[ip] = c
		}

		// Remove old timestamps (outside window)
		valid := c.timestamps[:0]
		for _, t := range c.timestamps {
			if now.Sub(t) < window {
				valid = append(valid, t)
			}
		}
		c.timestamps = valid

		// Check limit
		if len(c.timestamps) >= limit {
			mu.Unlock()
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		// Add current request
		c.timestamps = append(c.timestamps, now)
		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/ping", rateLimiter(http.HandlerFunc(pingHandler)))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
