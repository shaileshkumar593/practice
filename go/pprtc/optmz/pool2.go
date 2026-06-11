package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

var userPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("Creating new User")
		return &User{}
	},
}

func main() {
	u := userPool.Get().(*User)
	u.Name = "Shailesh"
	u.Age = 30

	fmt.Println(u)

	// Reset before putting back
	u.Name = ""
	u.Age = 0

	userPool.Put(u)
}

/*

Use it for:

	Buffers

	Temporary structs

	JSON encoders

	High allocation objects

Do NOT use it for:

	Database connections

	Long-lived objects

	Shared state
*/
