package main

import (
	"fmt"
	"time"
)

/*
func (t Time) AddDate(years int, months int, days int) Time
func (t Time) Add(d Duration) Time
*/

func main() {
	now := time.Now()
	fmt.Println("\n Today date is :", now)

	after := now.AddDate(1, 0, 0)
	fmt.Println("\nAdd 1 year ", after)

	after = now.AddDate(0, 1, 0)
	fmt.Println("\n Added a Month :", after)

	after = now.AddDate(0, 0, 1)
	fmt.Println("\n Added a Day :", after)

	after = now.AddDate(2, 2, 5)
	fmt.Println("Date  :", after)

	after = now.Add(10 * time.Hour)
	fmt.Println("\n Date :", after)

	after = now.Add(10 * time.Second)
	fmt.Println("\n New Date :", after)

}
