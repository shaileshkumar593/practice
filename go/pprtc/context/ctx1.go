// example-1-auto-cancel/main.go
// Description: use r.Context() in HTTP handlers so long-running work (DB calls, goroutines) cancels when the client disconnects.
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
	// dummy DB used to demonstrate QueryRowContext; no real data required.
	db, _ := sql.Open("sqlite3", ":memory:")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // server will cancel if client disconnects
		log.Println("handler: started")
		// simulate a DB query function that is context-aware
		res, err := slowDBQuery(ctx, db)
		if err != nil {
			log.Println("handler:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "result: %s\n", res)
	})
	log.Println("server: listening :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func slowDBQuery(ctx context.Context, db *sql.DB) (string, error) {
	// simulate slow pre-check
	select {
	case <-time.After(3 * time.Second):
		// proceed to query
	case <-ctx.Done():
		return "", ctx.Err()
	}
	// Example: use QueryRowContext in real DB drivers to be cancellable.
	row := db.QueryRowContext(ctx, "SELECT 'ok'")
	var s string
	if err := row.Scan(&s); err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return s, nil
}
