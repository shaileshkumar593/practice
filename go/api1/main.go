package main

import (
	"log"
	"net/http"
)

func main() {

	service := NewUserService()
	handler := NewHandler(service)

	http.HandleFunc("GET /users", handler.GetUsers)
	http.HandleFunc("POST /users", handler.CreateUser)
	http.HandleFunc("PUT /users", handler.UpdateUser)
	http.HandleFunc("DELETE /users", handler.DeleteUser)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
