package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

var db *gorm.DB

func main() {

	_db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = _db

	if err := db.AutoMigrate(&Grocery{}); err != nil {
		panic(err)
	}

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/allgroceries", AllGroceries)                        // ----> To request all groceries
	r.HandleFunc("/groceries/{name}", SingleGrocery)                   // ----> To request a specific grocery
	r.HandleFunc("/groceries", GroceriesToBuy).Methods("POST")         // ----> To add  new grocery to buy
	r.HandleFunc("/groceries/{name}", UpdateGrocery).Methods("PUT")    // ----> To update a grocery
	r.HandleFunc("/groceries/{name}", DeleteGrocery).Methods("DELETE") // ----> Delete a grocery
	log.Fatal(http.ListenAndServe(":10000", r))
}
