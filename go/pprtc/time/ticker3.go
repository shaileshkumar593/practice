package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Program Started")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	//-----------------------------------
	// 1. Fast ticker (1 sec)
	//-----------------------------------

	wg.Add(1)

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		counter := 0

		for {
			select {

			case t := <-ticker.C:
				counter++
				fmt.Printf("[FAST] Tick %d at %v\n", counter, t.Format("15:04:05"))

			case <-ctx.Done():
				fmt.Println("[FAST] Stopped")
				return
			}
		}
	}()

	//-----------------------------------
	// 2. Slow ticker (3 sec)
	//-----------------------------------

	wg.Add(1)

	go func() {

		defer wg.Done()

		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		for {

			select {

			case <-ticker.C:

				fmt.Println("[SLOW] Cache Refreshed")

			case <-ctx.Done():

				fmt.Println("[SLOW] Stopped")

				return
			}

		}

	}()

	//-----------------------------------
	// 3. Dynamic ticker replacement
	//-----------------------------------

	wg.Add(1)

	go func() {

		defer wg.Done()

		duration := 2 * time.Second

		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		count := 0

		for {

			select {

			case <-ticker.C:

				count++

				fmt.Println("[DYNAMIC] Tick", count)

				if count == 3 {

					fmt.Println("[DYNAMIC] Changing interval to 500ms")

					ticker.Stop()

					duration = 500 * time.Millisecond

					ticker = time.NewTicker(duration)
				}

			case <-ctx.Done():

				fmt.Println("[DYNAMIC] Stopped")

				ticker.Stop()

				return
			}
		}

	}()

	//-----------------------------------
	// Wait
	//-----------------------------------

	wg.Wait()

	fmt.Println("Program Finished")
}