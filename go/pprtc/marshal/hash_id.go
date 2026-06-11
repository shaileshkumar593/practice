package main

import (
	"fmt"
	"platform-svc-ofac/utils/secure"

	"time"
)

func UnixTimestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func GenerateRandomHash(hashLength int) string {
	unixTime := UnixTimestamp()
	return fmt.Sprintf(
		"%s%s",
		unixTime,
		secure.RandomString(hashLength-len(unixTime), []rune(secure.RuneAlNumCS)),
	)
}

func main() {
	fmt.Println("Hash code  ")
	fmt.Println("")
	fmt.Println()
}
