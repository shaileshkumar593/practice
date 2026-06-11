package main

import (
	"fmt"
	"runtime"
	"sync"
)

type MyObject struct {
	Data [1024]byte // A relatively large object
}

var objectPool = sync.Pool{
	New: func() interface{} {
		// This function is called when a new object is needed and none are available in the pool.
		return &MyObject{}
	},
}

func allocateDirectly() {
	_ = &MyObject{} // Allocates on heap
}

func allocateFromPool() {
	obj := objectPool.Get().(*MyObject) // Get object from pool
	// Do something with obj
	objectPool.Put(obj) // Return object to pool
}

func main() {
	// Let's observe memory before and after allocations
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Initial Alloc = %v Bytes\n", m.Alloc)

	// Simulate direct allocations
	for i := 0; i < 10000; i++ {
		allocateDirectly()
	}
	runtime.GC() // Force GC to see reclaimed memory
	runtime.ReadMemStats(&m)
	fmt.Printf("After Direct Allocations & GC: Alloc = %v Bytes\n", m.Alloc)

	// Simulate allocations from pool
	// Reset stats roughly (GC may clean up some previous direct allocations)
	runtime.GC() 
	runtime.ReadMemStats(&m)
	fmt.Printf("Before Pool Allocations: Alloc = %v Bytes\n", m.Alloc)
	for i := 0; i < 10000; i++ {
		allocateFromPool()
	}
	runtime.GC() // Force GC
	runtime.ReadMemStats(&m)
	fmt.Printf("After Pool Allocations & GC: Alloc = %v Bytes\n", m.Alloc)

	// You will observe that 'Alloc' increases much less or even stays stable
	// when using the pool compared to direct allocations, because objects are reused.
}