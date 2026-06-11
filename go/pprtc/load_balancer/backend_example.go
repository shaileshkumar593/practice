package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	port         = flag.Int("port", 8001, "Port to listen on")
	serverID     = flag.String("id", "server1", "Server identifier")
	requestCount atomic.Int64
)

func main() {
	flag.Parse()

	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/", handleRequest)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("[%s] Starting backend server on %s", *serverID, addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// handleHealth returns 200 OK for health checks
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok","server":"%s","requests":%d}`, *serverID, requestCount.Load())
}

// handleRequest processes regular requests
func handleRequest(w http.ResponseWriter, r *http.Request) {
	count := requestCount.Add(1)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Served-By", *serverID)

	// Simulate some processing
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{
  "server": "%s",
  "path": "%s",
  "method": "%s",
  "request_number": %d,
  "timestamp": "%s"
}`,
		*serverID,
		r.URL.Path,
		r.Method,
		count,
		time.Now().Format(time.RFC3339))

	log.Printf("[%s] %s %s -> %d (total: %d)", *serverID, r.Method, r.URL.Path, http.StatusOK, count)
}
