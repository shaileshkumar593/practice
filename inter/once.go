package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup

func initDB() {
	fmt.Println("DB Initialized")
}

func main() {

	for i := 0; i < 5; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			once.Do(initDB)
		}()
	}

	wg.Wait()
}
