package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	jobs := 10

	for i := 1; i <= jobs; i++ {
		jobID := i
		g.Go(func() error {
			fmt.Printf("Processing job %d\n", jobID)
			time.Sleep(500 * time.Millisecond)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("✅ All jobs done.")
}
