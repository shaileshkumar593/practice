package main

import (
	"fmt"
	"time"
)

func retry(attempts int, operation func() error) error {

	for i := 0; i < attempts; i++ {

		err := operation()

		if err == nil {
			return nil
		}

		fmt.Println("Retrying...")

		time.Sleep(time.Second)
	}

	return fmt.Errorf("failed after retries")
}

func main() {

	count := 0

	err := retry(3, func() error {

		count++

		if count < 3 {
			return fmt.Errorf("temporary failure")
		}

		fmt.Println("Success")

		return nil
	})

	fmt.Println(err)
}
