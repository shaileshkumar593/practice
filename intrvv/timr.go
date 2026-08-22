package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	for i := 0; i < 5; i++ {
		<-timer.C

		fmt.Println("Iteration:", i)

		timer.Reset(1 * time.Second)
	}
}
