package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Starting the news monitor at", time.Now().Format("15:04:05"))

	// Channel to receive news items from a separate goroutine
	newsFeed := make(chan string)

	// Start a simulated news producer
	go simulateNewsProducer(newsFeed)

	// Define the timeout duration
	timeoutDuration := 2 * time.Second

	// Create the timer once
	timer := time.NewTimer(timeoutDuration)

	select {
	case newsItem := <-newsFeed:
		// We received the news item before the timer expired.
		// It's crucial to stop the timer to free up resources immediately.
		if timer.Stop() {
			fmt.Printf("Received news: \"%s\" (Stopped timer)\n", newsItem)
		} else {
			// If Stop returns false, the timer has already expired or been stopped.
			fmt.Printf("Received news: \"%s\" (Timer already expired)\n", newsItem)
		}

	case <-timer.C:
		// The timer channel received a value, meaning the duration elapsed.
		// The newsFeed channel case was not selected in time.
		fmt.Printf("Timeout occurred after %v. No news received.\n", timeoutDuration)
	}

	fmt.Println("Finished monitoring at", time.Now().Format("15:04:05"))
}

// simulateNewsProducer is a helper function that sends a single piece of news
// after a random delay (0 to 3 seconds)
func simulateNewsProducer(c chan string) {
	// Seed the random generator for variable delays
	rand.Seed(time.Now().UnixNano())

	// Sleep for a random duration up to 3 seconds
	delay := time.Duration(rand.Intn(3000)) * time.Millisecond
	time.Sleep(delay)

	// Send the news item if the delay was short enough
	c <- fmt.Sprintf("Headline: Breaking news arrived after %v", delay)
}
