package main

import (
	"fmt"
	"time"
)

func main() {

	// Declaring t and u in UTC
	t := time.Date(2029, 0, 0, 0, 0, 0, 0, time.UTC)
	u := time.Date(2028, 0, 0, 0, 0, 0, 0, time.UTC)

	// Calling Before method
	res := t.Before(u)

	// Prints output
	fmt.Printf("%v", res)
	fmt.Println()

	fmt.Println(time.Now().Add(0).Format("2006-01-02 15:04:05"))
}
