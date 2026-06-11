package main

import (
	"fmt"
	"time"
)

func GetTimestampInMySQLAddBy(d time.Duration) string {
	return time.Now().Add(d).Format("2006-01-02 15:04:05")
}

func main() {
	expires := GetTimestampInMySQLAddBy(time.Hour * 24 * 7) // 7month
	fmt.Println("Time : ", expires)

	exp, _ := time.Parse(time.RFC3339, expires)

	if ok := time.Now().Before(exp); ok {
		fmt.Println("ok Before", ok)
	} else {
		fmt.Println("else Before:", ok)
	}

	if ok := time.Now().After(exp); ok {
		fmt.Println("ok After", ok)
	} else {
		fmt.Println("else: After ", ok)
	}
}
