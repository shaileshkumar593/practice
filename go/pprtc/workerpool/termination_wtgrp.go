package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	// Create a context that will be cancelled after 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Good practice to call cancel, though timeout does it too

	wg.Add(1)
	go func() {
		defer wg.Done() // Ensure Done is called

		// Create the request *with the context*. This is key!
		req, err := http.NewRequestWithContext(ctx, "GET", "http://httpbin.org/delay/10", nil) // 10 sec delay > 3 sec timeout
		if err != nil {
			// Note: Error checking context errors might be needed in real code
			fmt.Println("Error creating request:", err)
			return
		}

		fmt.Println("Sending request with context...")
		resp, err := http.DefaultClient.Do(req)
		// Now, if the context times out (after 3s), this Do() call will return an error (usually `context.DeadlineExceeded)
		if err != nil {
			fmt.Println("Error sending request:", err) // Will likely print context deadline exceeded
			return
		}
		// If we get here, the request finished *before* the timeout
		defer resp.Body.Close()
		fmt.Println("Request finished within timeout.")
	}()

	fmt.Println("Waiting with wg.Wait()...")
	// Now wg.Wait() is safe. Because the HTTP call respects the context,
	// the goroutine is guaranteed to finish (either successfully or via context error),
	// which means defer wg.Done() will eventually run. No leak!
	wg.Wait()
	fmt.Println("WaitGroup finished.")
}

/*
	So, main continues, but we've potentially left two goroutines hanging around, consuming resources, unable to finish. They are leaked. This is bad, especially in long-running server applications. Like I said earlier: always reason about how your goroutines end!

	The WaitGroup itself doesn't support cancellation. wg.Wait() is fundamentally a blocking call until the counter reaches zero. The real fix is to make sure the work being done inside the goroutines can be cancelled.
*/
