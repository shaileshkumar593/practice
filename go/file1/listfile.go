package main

import (
	"fmt"
	"os"
)

func main() {
	entries, err := os.ReadDir(".")

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}
