package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done() // Done is deferred correctly...
		// ...BUT this HTTP call has no timeout / context! It could block forever.
		req, err := http.NewRequest("GET", "http://httpbin.org/delay/10", nil) // 10 sec delay
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}
		fmt.Println("Sending request...")
		resp, err := http.DefaultClient.Do(req) // This blocks...
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("Request finished.") // ...potentially never reaching here
	}()

	// Wait for wg in a separate goroutine and signal via channel
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Use select to wait on the done channel OR a timeout
	select {
	case <-done:
		fmt.Println("WaitGroup finished normally.")
	case <-time.After(5 * time.Second): // Timeout after 5 seconds
		fmt.Println("Timeout waiting for WaitGroup!")
		// THE PROBLEM: We timed out, but the goroutines might still be running!
	}
	// Give the lingering goroutine a moment to print (or not)
	time.Sleep(1 * time.Second)
	fmt.Println("Main function exiting.")
}

/*
	deadlock avoided. But wait, there's more! What about timeouts and cancellation*/
