package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("data.txt")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File doesn't exist")
			return
		}

		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File exists")
	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size())
	fmt.Println("Directory:", info.IsDir())
}

/*
info.Name()
info.Size()
info.Mode()
info.ModTime()
info.IsDir()
*/
