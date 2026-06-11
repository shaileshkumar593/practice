package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// Intend to wait for one goroutine
	wg.Add(1)
	go func() {
		// PROBLEM: Defer is missing!
		req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			// We return early, wg.Done() is never called!
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			// We return early again, wg.Done() is never called!
			return
		}
		defer resp.Body.Close()
		// Only call Done() on the happy path
		wg.Done() // <<< If errors happen, this line is skipped!
	}()
	wg.Wait() // <<< This will wait FOREVER if an error occurs above.
}

/*
	In case of an error during request creation or sending, the Done() method is skipped.
	The WaitGroup counter never reaches zero, and our main goroutine calling wg.Wait() blocks indefinitely. Classic deadlock!*/
