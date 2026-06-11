package main

/*
	Problem:
Find the length of the longest substring without repeating characters.

	Time: O(n)
	Space: O(min(n, charset))
*/

import "fmt"

func lengthOfLongestSubstring(s string) (int, string) {
	m := make(map[byte]int)
	maxLen, left := 0, 0
	start := 0 // start index of longest substring

	for right := 0; right < len(s); right++ {
		// sliding window
		if pos, ok := m[s[right]]; ok && pos >= left {
			left = pos + 1
		}
		m[s[right]] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
			start = left
		}
	}

	// Extract the substring using start index and maxLen
	longestSubstring := s[start : start+maxLen]
	return maxLen, longestSubstring
}

func main() {
	s := "abcabcbbcdertyu"
	length, substring := lengthOfLongestSubstring(s)
	fmt.Printf("Length: %d, Substring: %s\n", length, substring) // Length: 3, Substring: abc
}

/*

	| Right | Char | Map (last index) | Left | Window | maxLen |
	| ----- | ---- | ---------------- | ---- | ------ | ------ |
	| 0     | a    | {a:0}            | 0    | a      | 1      |
	| 1     | b    | {a:0, b:1}       | 0    | ab     | 2      |
	| 2     | c    | {a:0, b:1, c:2}  | 0    | abc    | 3      |
	| 3     | a    | {a:3, b:1, c:2}  | 1    | bca    | 3      |
	| 4     | b    | {a:3, b:4, c:2}  | 2    | cab    | 3      |
	| 5     | c    | {a:3, b:4, c:5}  | 3    | abc    | 3      |
	| 6     | b    | {a:3, b:6, c:5}  | 5    | cb     | 3      |
	| 7     | b    | {a:3, b:7, c:5}  | 7    | b      | 3      |.

Start: right = 0, char = 'a'

1. [a] b c a b c b b
left=0, right=0, maxLen=1
right = 1, char = 'b'

2. [a b] c a b c b b
left=0, right=1, maxLen=2
right = 2, char = 'c'

3. [a b c] a b c b b
left=0, right=2, maxLen=3
right = 3, char = 'a' (repeat!)

Move left after previous 'a' (index 0 → 1)

4. b [c a] b c b b
left=1, right=3, maxLen=3
right = 4, char = 'b' (repeat!)

Move left after previous 'b' (index 1 → 2)

5. c [a b] c b b
left=2, right=4, maxLen=3
right = 5, char = 'c' (repeat!)

Move left after previous 'c' (index 2 → 3)

6. a [b c] b b
left=3, right=5, maxLen=3
right = 6, char = 'b' (repeat!)

Move left after previous 'b' (index 4 → 5)

7. c [b] b
left=5, right=6, maxLen=3
right = 7, char = 'b' (repeat!)

Move left after previous 'b' (index 6 → 7)

[b]
left=7, right=7, maxLen=3

*/
