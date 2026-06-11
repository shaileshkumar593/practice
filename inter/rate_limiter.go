package main

import (
	"fmt"
	"time"
)

func main() {

	limiter := time.Tick(time.Second)

	for i := 1; i <= 5; i++ {

		<-limiter

		fmt.Println("Request", i,
			time.Now().Format("15:04:05"))
	}
}
