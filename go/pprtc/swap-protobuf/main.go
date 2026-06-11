package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"

	"swap-protobuf/proto"
)

func main() {
	numbers := &proto.Numbers{A: 10, B: 20}
	fmt.Printf("Before Swap: A = %d, B = %d\n", numbers.A, numbers.B)

	data, err := proto.Marshal(numbers)
	if err != nil {
		log.Fatalf("Failed to encode: %v", err)
	}

	if err := os.WriteFile("numbers.bin", data, 0644); err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	readData, err := os.ReadFile("numbers.bin")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var decoded proto.Numbers
	if err := proto.Unmarshal(readData, &decoded); err != nil {
		log.Fatalf("Failed to decode: %v", err)
	}

	decoded.A, decoded.B = decoded.B, decoded.A
	fmt.Printf("After Swap:  A = %d, B = %d\n", decoded.A, decoded.B)
}
