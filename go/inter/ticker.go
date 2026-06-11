package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	defer ticker.Stop()

	request := []int{1, 2, 3, 4, 5}

	for _, req := range request {
		<-ticker.C
		fmt.Println("Processing request", req, time.Now())
	}
}
