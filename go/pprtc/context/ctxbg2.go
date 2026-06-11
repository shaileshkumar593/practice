package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("ctx.Err() : ", ctx.Err())
	fmt.Println("ctx.Done() : ", ctx.Done())
	fmt.Println("ctx.Value(\"key\") : ", ctx.Value("key"))
	fmt.Print("ctx.Deadline() : ")
	fmt.Print(ctx.Deadline())
	fmt.Println("")
}

/*
All these data are nil or empty currently because we have an
empty Context the Background context which is never canceled,
has no deadline and has no values. Background is typically used
in main, init, and tests and as the top-level Context for incoming
requests.*/
