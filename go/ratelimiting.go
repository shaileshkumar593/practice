package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	countrequest int
	windowEnd    time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)

	limit = 5

	window = 1 * time.Minute
)

func getIp(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func rateLimiter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := getIp(r)

		mu.Lock()

		client, exists := clients[ip]

		now := time.Now()

		if !exists || now.After(client.windowEnd) {
			// new window

			client = &Client{
				countrequest: 1,
				windowEnd:    now.Add(window),
			}
			clients[ip] = client
		} else {
			client.countrequest++
		}

		if client.countrequest > limit {
			mu.Unlock()
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		mu.Unlock()
		next(w, r)
	}
}

func pingHandlers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", rateLimiter(pingHandlers))

	fmt.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
