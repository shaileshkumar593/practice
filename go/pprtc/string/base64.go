package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encodedString := "VGhpcyBpcyBhIHRlc3Q="
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		fmt.Println("Error decoding string:", err)
		return
	}
	fmt.Println(string(decodedBytes))
}
