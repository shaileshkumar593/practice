package main

import (
	"fmt"
	"strings"
)

func FirstUnique(s string) {
	for i := 0; i < len(s); i++ {
		cnt := strings.Count(s, string(s[i]))

		if cnt == 1 {
			fmt.Println(string(s[i]), " : ", strings.Index(s, string(s[i])))
			break
		}
	}
}

func main() {
	input := "loveleetcode"

	FirstUnique(input)
}
