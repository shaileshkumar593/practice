package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	p := []rune(s)

	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}

	return string(p)
}

func main() {
	s := "welcome to sajjan pur"
	wordlist := strings.Fields(s)

	wordlist1 := make([]string, len(wordlist))

	for index, word := range wordlist {
		wordlist1[index] = Reverse(word)
	}

	fmt.Println(strings.Join(wordlist1, "--"))
}
