package main

import (
	"fmt"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	fmt.Fprintf(w, "User ID: %s", id)
}

func main() {

	http.HandleFunc("GET /users/{id}", getUser)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}