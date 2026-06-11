package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	<-time.After(2 * time.Second)

	fmt.Println("After 2 seconds")
}
