package main

import (
	"net/http"
)

func redirectToFreshman(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "https://freecodecamp.org/", http.StatusMovedPermanently)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/redirect", redirectToFreshman)

	http.ListenAndServe(":5000", mux)
}
