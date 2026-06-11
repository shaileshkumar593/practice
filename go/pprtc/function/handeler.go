package main

import (
	"fmt"
	"net/http"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received:", r.URL.Path)
		next(w, r)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", loggingMiddleware(helloHandler))
	http.ListenAndServe(":8080", nil)
}

/*
	It means:

		Takes a handler

		Returns a new handler

		So it:

		handler → wrapped handler
		This is a closure-based decorator pattern.
*/
