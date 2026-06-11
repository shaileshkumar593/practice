package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}
var ct int = 0
var m sync.RWMutex = sync.RWMutex{}

func hello() {

	fmt.Printf("Hello Go %v\n", ct)
	m.RUnlock()
	wg.Done()
}

func counter() {

	ct++
	m.Unlock()
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(11)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.RLock()
		go hello()
		m.Lock()
		go counter()
	}
	wg.Wait()
}
