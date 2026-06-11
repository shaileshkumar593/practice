package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// struct data
	type Addr struct {
		PostalCode int    `db:"postalcode"`
		Country    string `db:"country"`
	}
	type Me struct {
		Name    string      `db:"name"`
		Age     int         `db:"age"`
		Admin   bool        `db:"admin"`
		Hobbies []string    `db:"hobbies"`
		Address Addr        `db:"address"`
		Null    interface{} `db:"null"`
	}

	addr := Addr{
		PostalCode: 1111,
		Country:    "Japan",
	}
	me := Me{
		Name:    "noknow",
		Age:     2,
		Admin:   true,
		Hobbies: []string{"IT", "Travel"},
		Address: addr,
		Null:    nil,
	}

	// Convert map to json string
	jsonStr, err := json.Marshal(me)
	if err != nil {
		fmt.Println(err)
	}

	// Convert struct
	var mapData map[string]interface{}
	if err := json.Unmarshal(jsonStr, &mapData); err != nil {
		fmt.Println(err)
	}

	// Output
	fmt.Printf("Name: %v (%T)\n", mapData["Name"], mapData["Name"])
	fmt.Printf("Age: %v (%T)\n", mapData["Age"], mapData["Age"])
	fmt.Printf("Admin: %v (%T)\n", mapData["Admin"], mapData["Admin"])
	fmt.Printf("Hobbies: %v (%T)\n", mapData["Hobbies"], mapData["Hobbies"])
	fmt.Printf("Address: %v (%T)\n", mapData["Address"], mapData["Address"])
	fmt.Printf("Null: %v (%T)\n", mapData["Null"], mapData["Null"])

	delete(mapData, "Address")
	delete(mapData, "Null")
	fmt.Println("")
	fmt.Println("accessing element from map")
	for k, v := range mapData {
		fmt.Println(k, "----->", v)
	}
	fmt.Println("")
	fmt.Printf("Name: %v (%T)\n", mapData["Name"], mapData["Name"])
	fmt.Printf("Age: %v (%T)\n", mapData["Age"], mapData["Age"])
	fmt.Printf("Admin: %v (%T)\n", mapData["Admin"], mapData["Admin"])
	fmt.Printf("Hobbies: %v (%T)\n", mapData["Hobbies"], mapData["Hobbies"])
	fmt.Printf("Address: %v (%T)\n", mapData["Address"], mapData["Address"])
	fmt.Printf("Null: %v (%T)\n", mapData["Null"], mapData["Null"])
}

//https://hossainemruz.gitbook.io/notes/go/conversion/map-string-interface-to-struct
