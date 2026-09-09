package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler12(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from HTTP/2\n")

	fmt.Println("Protocol:", r.Proto)
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler12)

	log.Println("Server running on https://localhost:8443")

	err := http.ListenAndServe(
		":8443",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
}
