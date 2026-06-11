package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"newapi/model"
	"newapi/service"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		service: service,
	}
}

func (h *EmployeeHandler) EmployeeHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		employees := h.service.GetEmployees()

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(employees)

	case http.MethodPost:

		var emp model.Employee

		err := json.NewDecoder(r.Body).Decode(&emp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.service.CreateEmployee(emp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(emp)

	default:

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *EmployeeHandler) EmployeeByIDHandler(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/employees/")

	switch r.Method {

	case http.MethodGet:

		emp, exists := h.service.GetEmployee(id)

		if !exists {
			http.Error(w, "employee not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(emp)

	case http.MethodPut:

		var emp model.Employee

		err := json.NewDecoder(r.Body).Decode(&emp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		emp.ID = id

		err = h.service.UpdateEmployee(id, emp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(emp)

	case http.MethodDelete:

		err := h.service.DeleteEmployee(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Write([]byte("employee deleted"))

	default:

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
