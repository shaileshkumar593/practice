✅ 1. Goroutine + Channel (No sync package)
Simple worker pool using channels only
package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * job
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	// 5 workers
	for i := 1; i <= 5; i++ {
		go worker(i, jobs, results)
	}

	// Sending jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Receiving results
	for i := 1; i <= 10; i++ {
		fmt.Println(<-results)
	}
}
✅ 2. Goroutine + Channel + select
Worker pool with non-blocking select
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			results <- job * job
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	for i := 1; i <= 5; i++ {
		go worker(i, jobs, results)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println(<-results)
	}
}
✅ 3. sync.WaitGroup + Goroutine
Classic parallel computation with WaitGroup
package main

import (
	"fmt"
	"sync"
)

func square(n int, wg *sync.WaitGroup, out chan<- int) {
	defer wg.Done()
	out <- n * n
}

func main() {
	var wg sync.WaitGroup
	out := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go square(i, &wg, out)
	}

	wg.Wait()
	close(out)

	for v := range out {
		fmt.Println(v)
	}
}
✅ 4. Mutex + Goroutine
Safe shared counter
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final:", counter)
}
✅ 5. RWMutex + Goroutines
Read-heavy workload
package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[int]int
}

func main() {
	sm := SafeMap{m: make(map[int]int)}

	var wg sync.WaitGroup

	// Many readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sm.mu.RLock()
			fmt.Println("Reader", id, sm.m)
			sm.mu.RUnlock()
		}(i)
	}

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		sm.mu.Lock()
		sm.m[1] = 100
		sm.mu.Unlock()
	}()

	wg.Wait()
}
✅ 6. Atomic + Goroutine
High-performance lock-free counter
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Final:", atomic.LoadInt64(&counter))
}
✅ 7. sync.Once + Goroutine
Initialize service only once
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initDB() {
	fmt.Println("DB initialized")
}

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			once.Do(initDB)
		}()
	}

	fmt.Scanln()
}
✅ 8. Fan-in (Merge multiple channels into one)
package main

import (
	"fmt"
)

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)

	for _, ch := range channels {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
		}(ch)
	}
	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		a <- 10
		close(a)
	}()

	go func() {
		b <- 20
		close(b)
	}()

	out := fanIn(a, b)
	fmt.Println(<-out)
	fmt.Println(<-out)
}
✅ 9. Fan-out (Multiple workers reading same channel)
package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Println("Worker", id, "processed", job)
	}
}

func main() {
	jobs := make(chan int)

	for i := 1; i <= 5; i++ {
		go worker(i, jobs)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs)
}
✔️ 10. Channel + Mutex + Goroutines Together
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	counter := 0

	jobs := make(chan int)

	// 5 workers
	for i := 0; i < 5; i++ {
		go func() {
			for range jobs {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		jobs <- i
	}
	close(jobs)

	fmt.Println("Final:", counter)
}
