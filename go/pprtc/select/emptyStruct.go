package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	c := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			c <- struct{}{} // acquire lock
			defer func() {
				<-c // release lock
				wg.Done()
			}()

			mu.Lock()
			defer mu.Unlock()
			fmt.Printf("goroutine %d: hello\n", i)
		}(i)
	}

	wg.Wait()
}
