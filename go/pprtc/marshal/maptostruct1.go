package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// map data
	mapData := map[string]interface{}{
		"Name":    "noknow",
		"Age":     2,
		"Admin":   true,
		"Hobbies": []string{"IT", "Travel"},
		"Address": map[string]interface{}{
			"PostalCode": 1111,
			"Country":    "Japan",
		},
		"Null": nil,
	}

	// struct - Need to be defined according to the above map data.
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

	// Convert map to json string
	jsonStr, err := json.Marshal(mapData)
	if err != nil {
		fmt.Println(err)
	}

	// Convert json string to struct
	var me Me
	if err := json.Unmarshal(jsonStr, &me); err != nil {
		fmt.Println(err)
	}

	// Output
	fmt.Printf("Name: %s\n", me.Name)
	fmt.Printf("Age: %d\n", me.Age)
	fmt.Printf("Admin: %t\n", me.Admin)
	fmt.Printf("Hobbies: %v\n", me.Hobbies)
	fmt.Printf("Address: %v\n", me.Address)
	fmt.Printf("Null: %v\n", me.Null)
}
