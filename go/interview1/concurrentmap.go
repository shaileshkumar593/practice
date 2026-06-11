package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	m  map[any]any
	mu sync.Mutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[any]any),
	}
}

func (cm *ConcurrentMap) SetElement(key any, value any) {

	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.m[key] = value
}

func (cm *ConcurrentMap) GetElement(key any) any {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	val, exist := cm.m[key]
	if !exist {
		return fmt.Errorf("key %d donot exist", key.(int))
	}
	return val
}

func main() {

	cm := NewConcurrentMap()
	wg := sync.WaitGroup{}

	// Insert data
	for i := 0; i < 5; i++ {
		wg.Add(1)
		key := 100 + i
		val := 4444 * i

		go func(k, v int) {
			defer wg.Done()
			cm.SetElement(k, v)
		}(key, val)
	}

	wg.Wait()
	// Read data
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(a int) {
			defer wg.Done()

			val := cm.GetElement(100 + a)
			fmt.Println("Key:", 100+a, "Value:", val)

		}(i)
	}

	wg.Wait()

	fmt.Println("Done")
}
