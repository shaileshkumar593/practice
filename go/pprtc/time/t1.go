package main

import (
	"fmt"
	"time"
)

// use the time.Format() function to format a time.Time value into a string, and the time.Parse()
// function to parse a string into a time.Time value.

// time.LoadLocation() function to load a time zone, and then use the time.In() method to convert a time.Time

// time.Duration type, which represents a length of time. You can perform various calculations and operations on time.Duration values,
//  such as adding or subtracting them from time.Time values.

func main() {
	now := time.Now()
	someTime := time.Date(2023, time.April, 15, 10, 30, 0, 0, time.UTC)
	fmt.Println("Now:", now)
	fmt.Println("Some Time:", someTime)
	formattedTime := now.Format("2006-01-02 15:04:05")
	parsedTime, err := time.Parse("2006-01-02 15:04:05", formattedTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	fmt.Println("Formatted Time:", formattedTime)
	fmt.Println("Parsed Time:", parsedTime)

	// Load a specific time zone
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}
	nyTime := now.In(loc)
	fmt.Println("New York Time:", nyTime)

	// Working with time.Duration
	duration := 2 * time.Hour
	futureTime := now.Add(duration)
	fmt.Println("Future Time (after 2 hours):", futureTime)

	// Create a time value for April 15, 2023 at 10:30 AM in the UTC time zone
	someTime1 := time.Date(2023, time.April, 15, 10, 30, 0, 0, time.UTC)
	fmt.Println("Some Time:", someTime1)

	// Load the "America/New_York" time zone
	location, _ := time.LoadLocation("America/New_York")

	// Convert the current time to the New York time zone
	timeInNewYork1 := time.Now().In(location)
	fmt.Println("Current Time in New York:", timeInNewYork1)

	// reading system time zone
	systemLocation, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println("Error loading system location:", err)
	}
	fmt.Println(" system Location :", systemLocation)
	systemTime := time.Now().In(systemLocation)
	fmt.Println("System Time Zone Time:", systemTime)

	// timwe zone
	// Get the current time in the local system's time zone
	localTime := time.Now()

	// The local time zone is automatically detected and associated with localTime
	fmt.Println("Local Time:", localTime)

	// Use the Time.Zone() method to get the time zone name and offset
	zoneName, offset := localTime.Zone()

	fmt.Printf("Time Zone Name: %s\n", zoneName)
	fmt.Printf("Offset (in seconds east of UTC): %d\n", offset)

	// local to UTC  conversion
	// Get the current time in the system's local time zone
	localTime1 := time.Now()
	fmt.Println(" localTime ", localTime1)
	fmt.Println("Local time:", localTime1.Format(time.RFC3339))

	// Convert the time instant to UTC
	utcTime := localTime1.UTC()
	fmt.Println("UTC time:", utcTime.Format(time.RFC3339))

}

/*
	Output:time.Time.Year(), time.Time.Month(), time.Time.Day(): Retrieve the year, month, and day components of a time value.
	time.Time.Hour(), time.Time.Minute(), time.Time.Second(): Retrieve the hour, minute, and second components of a time value.
	time.Time.Sub(): Calculate the time difference between two time.Time values.
	time.Time.Equal(): Check if two time.Time values represent the same point in time.
	By leveraging these advanced time operations and calculations,
	you can build sophisticated time-based functionality in your Go applications,
	enabling features such as scheduling, event management, and time-series data processing.*/
