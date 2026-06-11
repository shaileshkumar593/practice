package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	<-time.After(2 * time.Second) // Blocks the main goroutine for 2 seconds
	fmt.Println("2 seconds later")
}
