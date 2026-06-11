package main

import "fmt"

/*
	Problem:
		Find the longest palindromic substring in a string.

*/

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, maxLen := 0, 0
	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		if right-left-1 > maxLen {
			start = left + 1
			maxLen = right - left - 1
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return s[start : start+maxLen]
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s)) // "bab" or "aba"
}

/*

	Complexity:

	Time: O((m+n) log(m+n))

	Space: O(m+n)

*/
