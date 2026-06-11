package main

// inserting records into a PostgreSQL database with Go's database/sql package
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "pgslx1005"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "connect-db"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// insert a row
	sqlStatement := `INSERT INTO emp (ename, sal, email) 
    VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, "Smith", 800, "smith@acme.com")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
}
