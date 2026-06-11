package main

import (
	"log"
	"net/http"

	"newapi/handler"
	"newapi/repository"
	"newapi/service"
)

func main() {

	repo := repository.NewEmployeeRepository()

	serviceLayer := service.NewEmployeeService(repo)

	handlerLayer := handler.NewEmployeeHandler(serviceLayer)

	http.HandleFunc("/employees", handlerLayer.EmployeeHandler)
	http.HandleFunc("/employees/", handlerLayer.EmployeeByIDHandler)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
