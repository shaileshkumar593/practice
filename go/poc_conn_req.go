package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var requestCount int64

func pingHandler(w http.ResponseWriter, r *http.Request) {
	count := atomic.AddInt64(&requestCount, 1)

	fmt.Printf("request number %d \n", count)

	time.Sleep(2 * time.Second)

	fmt.Fprintf(w, "PONG from request number %d ", count)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", pingHandler)

	err := http.ListenAndServe(":8080", mux)

	/*
		for{
			conn := listner.Accept()
			go serveConnection(conn)
		}*/
	if err != nil {
		panic(err)
	}

}
