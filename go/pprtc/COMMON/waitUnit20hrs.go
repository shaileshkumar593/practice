package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-time.After(20 * time.Hour):
		fmt.Println("Unblocked after 20 hours")
	}
}
