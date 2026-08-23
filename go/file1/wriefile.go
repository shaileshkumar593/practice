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

	data := []byte("Hello from Go\n")

	n, err := file.Write(data)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	fmt.Println("Bytes written:", n)
}
