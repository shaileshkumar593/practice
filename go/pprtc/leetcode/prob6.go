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

		Time: O(n^2)

		Space: O(1)

	| i | Center Type | Left | Right | Palindrome Found | Max Length Update? | `start`, `maxLen` After |
	| - | ----------- | ---- | ----- | ---------------- | ------------------ | ----------------------- |
	| 0 | Single      | 0    | 0     | "b"              | Yes (1>0)          | start=0, maxLen=1       |
	| 0 | Double      | 0    | 1     | ""               | No                 | start=0, maxLen=1       |
	| 1 | Single      | 1    | 1     | "a"              | No                 | start=0, maxLen=1       |
	| 1 | Double      | 1    | 2     | ""               | No                 | start=0, maxLen=1       |
	| 2 | Single      | 2    | 2     | "bab"            | Yes (3>1)          | start=1, maxLen=3       |
	| 2 | Double      | 2    | 3     | ""               | No                 | start=1, maxLen=3       |
	| 3 | Single      | 3    | 3     | "a"              | No                 | start=1, maxLen=3       |
	| 3 | Double      | 3    | 4     | ""               | No                 | start=1, maxLen=3       |
	| 4 | Single      | 4    | 4     | "d"              | No                 | start=1, maxLen=3       |
	| 4 | Double      | 4    | 5     | ""               | No                 | start=1, maxLen=3       |

*/
