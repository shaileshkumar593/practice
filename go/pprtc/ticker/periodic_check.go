package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go refreshCache(ctx)

	time.Sleep(12 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func refreshCache(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("🔄 refreshing cache from DB...")
		case <-ctx.Done():
			fmt.Println("cache refresh stopped:", ctx.Err())
			return
		}
	}
}

//  You must refresh an in-memory cache from a database every 10 seconds, but only while the service is running.
