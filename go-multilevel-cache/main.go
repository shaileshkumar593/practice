package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go-multilevel-cache/cache"
	"go-multilevel-cache/repository"
	"go-multilevel-cache/service"
)

func main() {
	localCache := cache.NewLocalCache(10_000)

	redisCache := cache.NewRedisCache(
		"localhost:6379",
		"",
	)

	if err := redisCache.Ping(); err != nil {
		log.Printf("warning: Redis is unavailable: %v", err)
	}

	multiCache := cache.NewMultiLevelCache(
		localCache,
		redisCache,
		30*time.Second, // L1 TTL
		5*time.Minute,  // L2 TTL
	)

	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo, multiCache)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idText := strings.TrimPrefix(r.URL.Path, "/users/")
		id, err := strconv.Atoi(idText)
		if err != nil || id <= 0 {
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}

		user, source, err := userService.GetUser(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"source": source,
			"user":   user,
		})
	})

	mux.HandleFunc("/users/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user repository.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if user.ID <= 0 {
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}

		if err := userService.UpdateUser(r.Context(), user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "user updated and cache invalidated",
			"user":    user,
		})
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("API listening on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
