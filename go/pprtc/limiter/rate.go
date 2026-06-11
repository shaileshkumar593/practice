package main

import (
	"log"
	"net/http"
	"time"
)

var limiter = time.NewTicker(time.Second / 100) // 100 req/sec

func handler(w http.ResponseWriter, r *http.Request) {
	<-limiter.C // wait for a tick
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}

func main() {
	defer limiter.Stop()

	http.HandleFunc("/", handler)

	log.Println("Server running on :8080 (100 req/sec)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
