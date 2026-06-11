package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // Add before go
	go func() {
		defer wg.Done() // Defer immediately! Now it *always* runs on exit.
		req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return // Done() will still run thanks to defer
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return // Done() will still run
		}
		defer resp.Body.Close()
		// No need for wg.Done() here anymore
		fmt.Println("Request successful (in theory)!")
	}()
	wg.Wait()
	fmt.Println("Main goroutine finished waiting.")
}
