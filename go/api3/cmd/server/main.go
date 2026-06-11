package main

import (
	"log"
	"net/http"

	"api3/internal/handler"
	"api3/internal/service"
)

func main() {
	userService := service.NewUserService()

	userHandler := handler.NewUserHandler(
		userService,
	)

	mux := http.NewServeMux()

	mux.HandleFunc(
		"GET /users",
		userHandler.GetUsers,
	)

	mux.HandleFunc(
		"POST /users",
		userHandler.CreateUser,
	)

	mux.HandleFunc(
		"PUT /users",
		userHandler.UpdateUser,
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server started on :8080")

	log.Fatal(server.ListenAndServe())
}
