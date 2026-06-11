package main

/*
	In Go,

time.Sleep(d time.Duration)
expects an argument of type time.Duration, not an int.

So when you do:

time.Sleep(1 * time.Second)
✅ works fine because:

time.Second is already a time.Duration constant, internally 1e9 nanoseconds.

1 * time.Second → type time.Duration.

But when you do:

val := 5
time.Sleep(val * time.Second)
❌ you get a compile error like:

invalid operation: val * time.Second (mismatched types int and time.Duration)
Because:

val is an int

time.Second is a time.Duration

# Go does not auto-convert between int and time.Duration

✅ Correct Way
Convert val explicitly to time.Duration:

val := 5
time.Sleep(time.Duration(val) * time.Second)
Now it works perfectly.

🧠 Deep Dive
Let’s print types for clarity:

fmt.Printf("Type of val: %T\n", val)          // int
fmt.Printf("Type of time.Second: %T\n", time.Second) // time.Duration
So you must explicitly convert int → time.Duration.

🏁 Final Correct Example
*/

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Sleeping for %d seconds...\n", i)
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("Woke up!")
	}
}

/*
	✅ Output:

	Sleeping for 1 seconds...
	Woke up!
	Sleeping for 2 seconds...
	Woke up!
	Sleeping for 3 seconds...
	Woke up!
*/
