package main

import (
	"fmt"
	"sync"
)

func main() {

	// Uncomment ONE case at a time

	// case1()
	// case2()
	// case3()
	// case4()
	// case5()
	// case6()
}

/////////////////////////////////////////////////////
// CASE 1: Send on unbuffered channel without receiver
/////////////////////////////////////////////////////

func case1() {
	ch := make(chan int)
	ch <- 10 // ❌ No receiver
}

/////////////////////////////////////////////////////
// CASE 2: Receive on channel without sender
/////////////////////////////////////////////////////

func case2() {
	ch := make(chan int)
	<-ch // ❌ No sender
}

/////////////////////////////////////////////////////
// CASE 3: Circular deadlock between goroutines
/////////////////////////////////////////////////////

func case3() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		<-ch2
	}()

	go func() {
		ch2 <- 2
		<-ch1
	}()

	select {} // block main
}

/////////////////////////////////////////////////////
// CASE 4: WaitGroup deadlock (Done not called)
/////////////////////////////////////////////////////

func case4() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("Working...")
		// ❌ Forgot wg.Done()
	}()

	wg.Wait() // waits forever
}

/////////////////////////////////////////////////////
// CASE 5: Mutex deadlock (double lock)
/////////////////////////////////////////////////////

func case5() {
	var mu sync.Mutex

	mu.Lock()
	mu.Lock() // ❌ Deadlock (same goroutine)

	fmt.Println("Never reached")
}

/////////////////////////////////////////////////////
// CASE 6: Channel range without close
/////////////////////////////////////////////////////

func case6() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		// ❌ Channel never closed
	}()

	for v := range ch { // waits forever
		fmt.Println(v)
	}
}
