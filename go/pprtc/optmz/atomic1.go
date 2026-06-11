package main

import (
	"fmt"
	"sync/atomic"
)

// Atomic Load & Store (Safe Read / Write)

var flag uint32

func main() {

	// Store value safely
	atomic.StoreUint32(&flag, 1)

	// Load value safely
	if atomic.LoadUint32(&flag) == 1 {
		fmt.Println("Flag is set")
	}
}

/*
	Used for:

			Feature flags

			Circuit breaker state

			Shutdown signals
*/
