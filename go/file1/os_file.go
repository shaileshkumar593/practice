package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 10)

	for {
		n, err := file.Read(buffer)

		if n > 0 {
			fmt.Print(string(buffer[:n]))
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
	}
}
