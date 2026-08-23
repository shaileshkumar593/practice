package main

import (
	"fmt"
	"sync"
)

func process(id int) error {
	if id == 2 {
		return fmt.Errorf("job %d failed", id)
	}

	fmt.Printf("job %d succeeded\n", id)
	return nil
}

func main() {
	var wg sync.WaitGroup
	errCh := make(chan error, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			if err := process(id); err != nil {
				errCh <- err
			}
		}(i)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		fmt.Println("Error:", err)
	}
}
