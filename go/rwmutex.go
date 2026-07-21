package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.RWMutex
	value int
}

func (c *Counter) Read(id int) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	fmt.Printf("Reader %d reading value: %d\n", id, c.value)

	time.Sleep(2 * time.Second)

	fmt.Printf("Reader %d finished\n", id)
}

func (c *Counter) Write(val int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Println("Writer started")

	c.value = val

	time.Sleep(3 * time.Second)

	fmt.Println("Writer finished")
}

func main() {

	counter := &Counter{}

	var wg sync.WaitGroup

	// Three readers
	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			counter.Read(id)
		}(i)
	}

	time.Sleep(500 * time.Millisecond)

	// One writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		counter.Write(100)
	}()

	wg.Wait()
}