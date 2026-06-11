package main

import (
	"fmt"
	"time"
)

func main() {
	defer trackTime(time.Now())
	time.Sleep(time.Second)
}

func trackTime(t time.Time) {
	elapsedtime := time.Since(t)
	fmt.Printf("execution completed in %v\n", elapsedtime)
}
