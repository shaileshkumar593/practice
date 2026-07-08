package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	for {

		select {

		case t := <-timer.C:
			fmt.Println("Timer fired at:", t)
			return

		default:
			fmt.Println("Doing some work...")
			time.Sleep(time.Second)
		}
	}
}