package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int) {
	fmt.Println(id, "starting")
	time.Sleep(time.Second)
	fmt.Println(id, "ending")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		// using new Go way to call goroutine
		wg.Go(func() { Worker(i) })
	}

	wg.Wait()
	fmt.Println("All worker finish")
}
