package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	cutoff := now.Add(-1 * time.Minute)

	fmt.Println(now.After(cutoff)) // true
	fmt.Println(cutoff.After(now)) // false
}

/*
	Why?
If:

now     = 10:05
cutoff  = 10:04
Then:

10:05.After(10:04) → true
10:04.After(10:05) → false
*/
