package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			log.Println("Hey!")
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
}
