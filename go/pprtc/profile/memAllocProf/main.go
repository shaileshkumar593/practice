// main.go

package main

import "sync"

var pool = sync.Pool{

	New: func() interface{} {

		return new(int)

	},
}

func Add(a, b int) int {

	result := a + b

	val := pool.Get().(*int)

	*val = result

	defer pool.Put(val)

	return *val

}
