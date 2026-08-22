package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var rqcnt uint64

func handler1(w http.ResponseWriter, r *http.Request) {
	count := atomic.AddUint64(&rqcnt, 1)

	fmt.Fprintln(w, " request cnt ", count,
		" request method ", r.Method,
		" request path ", r.URL.Path)
}

func maiin() {
	http.HandleFunc("/get/", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
