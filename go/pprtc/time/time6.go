package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Today:", now)

	after := now.AddDate(-1, 0, 0)
	fmt.Println("Subtract 1 Year:", after)

	after = now.AddDate(0, -1, 0)
	fmt.Println("Subtract 1 Month:", after)

	after = now.AddDate(0, 0, -1)
	fmt.Println("Subtract 1 Day:", after)

	after = now.AddDate(-2, -2, -5)
	fmt.Println("Subtract multiple values:", after)

	after = now.Add(-10 * time.Minute)
	fmt.Println("Subtract 10 Minutes:", after)

	after = now.Add(-10 * time.Second)
	fmt.Println("Subtract 10 Second:", after)

	after = now.Add(-10 * time.Hour)
	fmt.Println("Subtract 10 Hour:", after)

	after = now.Add(-10 * time.Millisecond)
	fmt.Println("Subtract 10 Millisecond:", after)

	after = now.Add(-10 * time.Microsecond)
	fmt.Println("Subtract 10 Microsecond:", after)

	after = now.Add(-10 * time.Nanosecond)
	fmt.Println("Subtract 10 Nanosecond:", after)
}
