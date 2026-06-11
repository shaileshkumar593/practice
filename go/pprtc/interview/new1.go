package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go fmt.Println(<-ch)
		ch <- i
	}
	time.Sleep(5 * time.Second)
}

/*

	| Scenario                                             | Code Pattern                                                     | Channel Type      | What Happens                                                                              | Deadlock? | Key Reason                          |
| ---------------------------------------------------- | ---------------------------------------------------------------- | ----------------- | ----------------------------------------------------------------------------------------- | --------- | ----------------------------------- |
| **1. Buffered channel (works, random output)**       | `ch := make(chan int, 1)`<br>`ch <- i`<br>`go fmt.Println(<-ch)` | Buffered (size=1) | Send succeeds (buffer has space)<br>Receive happens in main<br>Print happens in goroutine | ❌ No      | Buffer prevents blocking            |
| **2. Unbuffered (send first → deadlock)**            | `ch := make(chan int)`<br>`ch <- i`<br>`go fmt.Println(<-ch)`    | Unbuffered        | Send blocks immediately<br>No receiver yet                                                | ✅ Yes     | No receiver available               |
| **3. Unbuffered (goroutine first → still deadlock)** | `go fmt.Println(<-ch)`<br>`ch <- i`                              | Unbuffered        | `<-ch` evaluated in main<br>Main blocks before send                                       | ✅ Yes     | Argument evaluated before goroutine |
*/
