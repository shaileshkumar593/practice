package main

import (
	"log"
	"net/http"
	"time"
)

var limiter = time.NewTicker(time.Second / 100) // 100 req/sec

func handler(w http.ResponseWriter, r *http.Request) {
	// Stop waiting if the client disconnects
	select {
	case <-limiter.C:
		// allowed
	case <-r.Context().Done():
		// client canceled the request
		return
	}

	// Response handling
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("OK\n")); err != nil {
		log.Println("write error:", err)
	}
}

func main() {
	defer limiter.Stop()

	http.HandleFunc("/", handler)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Server running on :8080 (100 req/sec)")
	log.Fatal(server.ListenAndServe())
}
