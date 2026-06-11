package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2015, 02, 21, 23, 10, 52, 211, time.UTC)
	fmt.Println("time is .....", t)

	fmt.Println("--------------------------------------------")
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	nanoseconds := t.Nanosecond()

	fmt.Println("year :", year)
	fmt.Println("Month :", month)
	fmt.Println("day :", day)
	fmt.Println("hour :", hour)
	fmt.Println("minute :", minute)
	fmt.Println("second :", second)
	fmt.Println("nanoseconds : ", nanoseconds)
}
