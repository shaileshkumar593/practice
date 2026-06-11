package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.RedirectHandler("https://freshman.tech", http.StatusMovedPermanently))

	http.ListenAndServe(":5000", mux)
}
