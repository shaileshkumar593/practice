package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	fmt.Println("File created successfully")
}

/*
	If the file doesn't exist, Go creates it.

If it already exists, os.Create truncates it, meaning its previous contents are removed.

*/
