package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "UserID", 123)

	go performTask34(ctx)

	// Continue with other operations
}

func performTask34(ctx context.Context) {
	userID := ctx.Value("UserID")
	fmt.Println("User ID:", userID)
}

/*
	we create a parent context using context.Background(). We then use context.WithValue() to attach a user ID to the context.
	 The context is then passed to the performTask goroutine, which retrieves the user ID using ctx.Value().*/
