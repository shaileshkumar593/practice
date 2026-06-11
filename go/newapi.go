package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{
		"status": "ok",
	}

	json.NewEncoder(w).Encode(resp)
}

// POST /decode
func decodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request

	// decode JSON body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "decoded successfully",
		"data":    req,
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthCheck)
	mux.HandleFunc("/decode", decodeHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)
}
