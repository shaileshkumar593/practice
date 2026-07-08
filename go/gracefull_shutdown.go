package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	go func() {
		fmt.Println("Server started on :8080")
		server.ListenAndServe()
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	<-ctx.Done()

	fmt.Println("Shutting down...")

	server.Shutdown(context.Background())

	fmt.Println("Server stopped")
}

/*
	signal.NotifyContext creates a context that is automatically canceled when one of the specified OS signals (such as os.Interrupt) is received.

	ctx.Done() returns a channel that closes when the context is canceled.

	<-ctx.Done() blocks until cancellation occurs.

	defer stop() unregisters the signal handler and releases associated resources.

	This pattern is the standard way to implement graceful shutdown in Go servers and long-running applications.


*/