package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Now().UTC()
	now := time.Now()
	zone, offset1 := now.Zone()
	// offset in secon diff CUR-UTC
	fmt.Println("timezone -------------------------- ", zone, offset1/3600)
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Println("Current date and time (RFC3339):", now.Format(time.RFC3339))
	fmt.Println("Current from date :", now.Format("2006-01-02 00:00:00"))
	fmt.Println("Current to date :", tomorrow.Format("2006-01-02 00:00:00"))
	fmt.Println("Current date and time (Mon Jan 2 15:04:05 MST 2006):", now.Format("Mon Jan 2 15:04:05 MST 2006"))
	dd := now.Format("02-01-2006")
	fmt.Println("--------------------------------", dd)

}
