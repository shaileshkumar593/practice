package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	done := ctx.Done()
	for i := 0; ; i++ {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			fmt.Println("tick", i)
		}
	}
}

//We can see that, in the context of the example, the loop goes
// on infinitely because the context is never completed.
