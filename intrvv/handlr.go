package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var requestcounter int64

func handler(w http.ResponseWriter, r *http.Request) {
	count := atomic.AddInt64(&requestcounter, 1)
	log.Println("Path:", r.URL.Path, "Count:", count)

	fmt.Fprintln(w, "Hello")
	fmt.Fprintln(w, "Request count is", count)
}

func main() {
	http.HandleFunc("/get/", handler) // callback function for handler
	log.Fatal(http.ListenAndServe(":8080", nil))

}
