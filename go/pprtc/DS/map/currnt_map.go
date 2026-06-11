package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Store(fmt.Sprintf("key%d", i), i)
		}(i)
	}

	wg.Wait()

	sm.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
