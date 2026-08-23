package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data written")
}
