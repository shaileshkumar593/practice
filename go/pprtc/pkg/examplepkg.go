package main

import (
	"fmt"
	e "pkg/example"
)

func main() {
	fmt.Println("In Main")
	fmt.Println("Example Package", e.Age())
}
