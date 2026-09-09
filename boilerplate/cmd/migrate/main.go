package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	command := flag.String("command", "up", "Migration command: up, down, version")
	flag.Parse()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/appdb?sslmode=disable"
	}

	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		log.Fatalf("failed to initialize migration: %v", err)
	}

	defer func() {
		sourceErr, databaseErr := m.Close()

		if sourceErr != nil {
			log.Printf("migration source close error: %v", sourceErr)
		}

		if databaseErr != nil {
			log.Printf("database close error: %v", databaseErr)
		}
	}()

	switch *command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migration up failed: %v", err)
		}

		fmt.Println("Migration completed successfully")

	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatalf("migration down failed: %v", err)
		}

		fmt.Println("Rollback completed successfully")

	case "version":
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatalf("failed to get migration version: %v", err)
		}

		fmt.Printf("Version: %d, Dirty: %v\n", version, dirty)

	default:
		log.Fatalf("unknown command: %s", *command)
	}
}
