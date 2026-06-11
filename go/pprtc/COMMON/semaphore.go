package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var m sync.Mutex
var i int = 0

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incrementtt()
	}
	wg.Wait()
}

func incrementtt() {
	m.Lock()
	i = i + 1
	fmt.Println(i)
	m.Unlock()
	wg.Done()
}

//go run -race filename
/*
 go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
$ go build -race mycmd   // to build the command
$ go install -race mypkg // to install the package
*/
