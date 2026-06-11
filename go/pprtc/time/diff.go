package main

import (
	"fmt"
	"time"
)

func GetTimeDifference(date string) (diff time.Duration) {
	old_date, _ := time.Parse(time.RFC3339, date)
	diff = time.Now().Sub(old_date)
	return
}
func main() {
	firstDate := time.Date(2022, 4, 13, 3, 0, 0, 0, time.UTC)
	secondDate := time.Date(2010, 2, 12, 6, 0, 0, 0, time.UTC)
	difference := GetTimeDifference(firstDate.UTC().Format("2006-01-02 15:04:05"))
	diff := GetTimeDifference(secondDate.UTC().Format("2006-01-02 15:04:05"))

	fmt.Printf("Years: %d\n", int64(diff.Hours()/24/365))
	fmt.Println("------------------", difference)
	fmt.Println("The difference between dates", firstDate, "and", secondDate, "is: ")
	fmt.Printf("Years: %d\n", int64(difference.Hours()/24/365))
	fmt.Printf("Months: %d\n", int64(difference.Hours()/24/30))
	fmt.Printf("Weeks: %d\n", int64(difference.Hours()/24/7))
	fmt.Printf("Days: %d\n", int64(difference.Hours()/24))
	fmt.Printf("Hours: %.f\n", difference.Hours())
	fmt.Printf("Minutes: %.f\n", difference.Minutes())
	fmt.Printf("Seconds: %.f\n", difference.Seconds())
	fmt.Printf("Nanoseconds: %d\n", difference.Nanoseconds())
}
