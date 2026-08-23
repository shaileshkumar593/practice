package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join(
		"data",
		"logs",
		"app.log",
	)

	fmt.Println(path)
}
