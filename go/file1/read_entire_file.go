package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println(string(data))
}
