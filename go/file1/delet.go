package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("data.txt")

	if err != nil {
		fmt.Println("Delete error:", err)
		return
	}

	fmt.Println("File deleted")
}
