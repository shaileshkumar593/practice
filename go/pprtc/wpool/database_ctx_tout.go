package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := sql.Open("postgres", "postgres://username:password@localhost/mydatabase?sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Execute the database query with the custom context
	rows, err := db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Process query results
}

/*
	In this example, we create a context with a timeout of 2 seconds using context.WithTimeout().
	We then open a connection to a PostgreSQL database using the sql.Open() function. When executing the database query with db.QueryContext(),
	the context ensures that the operation will be canceled if it exceeds the specified timeout.*/
