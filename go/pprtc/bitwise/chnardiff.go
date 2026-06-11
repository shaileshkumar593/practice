package main

import "fmt"

func findTheDifferenceMap(s, t string) byte {
	freq := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		freq[t[i]]--
		if freq[t[i]] < 0 {
			return t[i]
		}
	}
	return 0
}

func main() {
	fmt.Println(string(findTheDifferenceMap("abcd", "ceadb")))
}
