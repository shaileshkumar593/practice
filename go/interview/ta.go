package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		7*time.Second,
	)

	defer cancel()

	select {

	case <-time.After(5 * time.Second):

		fmt.Println("Completed")

	case <-ctx.Done():

		fmt.Println(ctx.Err())
	}
}
