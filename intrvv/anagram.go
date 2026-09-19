package main

import (
	"fmt"
)

/*
another example - listen and silent

input - eat and tea , output true
input - tic and toe - output false

*/

func FindAnagram(given, target string) bool {

	m := make(map[rune]int)
	t := make(map[rune]int)

	if len(given) != len(target) {
		return false
	}
	for _, val := range given {
		if intval, ok := m[val]; ok {
			m[val] = intval + 1
		} else {
			m[val] = 1
		}

	}

	for _, val := range target {
		if intval, ok := m[val]; ok {
			t[val] = intval + 1
		} else {
			t[val] = 1
		}

	}

	fmt.Print(t, m)
	for key, _ := range m {
		if m[key] != t[key] {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(FindAnagram("xyx", "xyy"))
}
