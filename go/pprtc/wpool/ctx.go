package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup a simple DB (can be any driver)
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Request context (auto-cancelled if client disconnects)
		ctx := r.Context()

		log.Println("Request started")

		// Simulate database work
		result, err := runQuery(ctx, db)
		if err != nil {
			// If client cancelled → ctx.Err() will be context.Canceled
			log.Println("Query stopped:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// If query finished before cancellation
		log.Println("Query completed successfully")
		fmt.Fprintf(w, "Result: %s\n", result)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func runQuery(ctx context.Context, db *sql.DB) (string, error) {

	// If you use QueryContext, ExecContext, etc, they obey ctx cancellation
	query := "SELECT 'Hello from DB'"

	// Simulate a slow DB query
	time.Sleep(3 * time.Second)

	// Check if context is already cancelled before firing
	select {
	case <-ctx.Done():
		// Request ended → stop now
		return "", ctx.Err()
	default:
	}

	// Now execute the DB query with context
	row := db.QueryRowContext(ctx, query)

	var result string
	err := row.Scan(&result)
	if err != nil {
		return "", err
	}

	return result, nil
}
