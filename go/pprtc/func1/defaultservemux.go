package main

import (
	"fmt"
	"net/http"
)

func logger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Before Handler")

		next(w, r)

		fmt.Println("After Handler")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Handler")
}

func main() {

	http.HandleFunc("/", logger(home))

	http.ListenAndServe(":8080", nil)
}
