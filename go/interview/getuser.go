package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	users := []User{
		{1, "Alice"},
		{2, "Bob"},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func main() {

	http.HandleFunc("/users", getUsers)

	http.ListenAndServe(":8080", nil)
}
