package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go pollAPI(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func pollAPI(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("📡 polling API...")
		case <-ctx.Done():
			fmt.Println("polling stopped:", ctx.Err())
			return
		}
	}
}

// You want to poll an external API but not exceed limits, e.g. once every 500ms.
