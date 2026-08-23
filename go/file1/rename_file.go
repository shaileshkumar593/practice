package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename(
		"data.txt",
		"newdata.txt",
	)

	if err != nil {
		fmt.Println("Rename error:", err)
		return
	}

	fmt.Println("File renamed")
}
