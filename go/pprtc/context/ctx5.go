// example-5-values/main.go
// Description: attach small request-scoped data (IDs, tokens). Do not store large payloads or business objects.
package main

import (
	"context"
	"fmt"
)

type ctxKey string

const traceKey ctxKey = "traceID"

func main() {
	ctx := context.WithValue(context.Background(), traceKey, "trace-1234")
	process(ctx)
}

func process(ctx context.Context) {
	if v := ctx.Value(traceKey); v != nil {
		fmt.Println("trace id:", v)
	} else {
		fmt.Println("no trace")
	}
}
