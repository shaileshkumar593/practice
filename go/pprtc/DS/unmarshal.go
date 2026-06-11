package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := `{"name": "Rob", "age": 18}`

	var obj any
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// obj now contains the parsed JSON data, but its type is "any"
	// We can use a type assertion to access its properties:
	m := obj.(map[string]any)
	fmt.Println("Name:", m["name"])
	fmt.Println("Age:", m["age"])
}
