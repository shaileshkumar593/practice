package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {

	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

}

/*
Going through the code, there are some similarities among context.
WithDeadline and context.WithTimeout, the difference here is that you
have to pass in a duration of type time.Duration when using WithTimeout
to end the Context. Our output is context deadline exceeded because we
set the timeout to 1millisecond and the time.After in our select block
will run for 1second.*/
