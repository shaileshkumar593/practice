package main

import (
	"fmt"
	"net/http"

	"go-api-project/handlers"
)

func main() {

	http.HandleFunc("/users", handlers.UserHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
