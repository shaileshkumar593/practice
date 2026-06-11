package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			log.Printf("Tick at: %v\n", t.UTC())
			// do something
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
}
