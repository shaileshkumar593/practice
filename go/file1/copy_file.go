package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	source, err := os.Open("source.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer source.Close()

	destination, err := os.Create("copy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)

	if err != nil {
		fmt.Println("Copy error:", err)
		return
	}

	fmt.Println("File copied")
}
