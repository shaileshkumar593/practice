package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	m  map[any]any
	mu sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[any]any, 0),
	}
}

func (cm *ConcurrentMap) SetElement(key any, value any) error {

	err := checktype(key)
	if err != nil {
		return err
	}
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

func checktype(v any) error {
	switch val := v.(type) {
	case int:
		return nil
	case string:
		return nil

	case [1]int:
		return nil

	case float:
		return nil
	case []int, []float, []string:
		return fmt.Errorf("slice is not use ")
	}
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
			err := cm.SetElement(k, v)
			fmt.Printf(err.errors())
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
