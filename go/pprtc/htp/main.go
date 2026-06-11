package main

import (
	"log"
	"net/http"
)

func redirectToTls(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://freshman.tech", http.StatusMovedPermanently)
}

func main() {
	if err := http.ListenAndServe(":5000", http.HandlerFunc(redirectToTls)); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
