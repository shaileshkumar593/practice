package main

import (
	"fmt"
	"os"
)

func main() {

	// 1. Create
	file, err := os.Create("user.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2. Write
	_, err = file.WriteString("John\n")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	file.Close()

	// 3. Read
	data, err := os.ReadFile("user.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Content:")
	fmt.Println(string(data))

	// 4. Append
	file, err = os.OpenFile(
		"user.txt",
		os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.WriteString("Alex\n")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	file.Close()

	// 5. Read again
	data, err = os.ReadFile("user.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("After append:")
	fmt.Println(string(data))

	// 6. Rename
	err = os.Rename("user.txt", "users.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 7. Delete
	err = os.Remove("users.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File deleted")
}
