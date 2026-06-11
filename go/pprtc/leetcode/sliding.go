package main


// longest unique substring contineous 
func longestSubstringStrSet(s string) string {
	set := make(map[byte]struct{})

	left := 0
	maxLen := 0
	start := 0

	for right := 0; right < len(s); right++ {

		// If duplicate, shrink window
		// This loop only runs when duplicates occur
👉 		// But every deletion corresponds to a previous insertion
		for {
			if _, exists := set[s[right]]; !exists {
				break
			}
			delete(set, s[left])
			left++
		}

		// Add current character
		set[s[right]] = struct{}{}

		// Update result
		if right-left+1 > maxLen {
			maxLen = right - left + 1
			start = left
		}
	}

	return s[start : start+maxLen]
}

// Inserted into set once

// Deleted from set once

func longestSubstringStr1(s string) string {
	index := make(map[byte]int)
	left := 0
	maxLen := 0
	start := 0

	for right := 0; right < len(s); right++ {
		if val, ok := index[s[right]]; ok && val >= left {
			left = val + 1
		}

		index[s[right]] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
			start = left
		}
	}

	return s[start : start+maxLen]
}

/*
	| Approach        | Inner Loop | Time           | Notes           |
| --------------- | ---------- | -------------- | --------------- |
| Index map       | No loop    | O(n)           | Faster          |
| Set (your code) | Yes        | O(n) amortized | Slightly slower |
*/


// Longest Substring with At Most K Distinct

func longestKDistinct(s string, k int) int {
	freq := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		for len(freq) > k {
			freq[s[left]]--
			if freq[s[left]] == 0 {
				delete(freq, s[left])
			}
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

/*

	| Step | Window | freq          | len(freq) | Action   |
| ---- | ------ | ------------- | --------- | -------- |
| e    | e      | {e:1}         | 1         | valid    |
| c    | ec     | {e:1,c:1}     | 2         | valid    |
| e    | ece    | {e:2,c:1}     | 2         | valid    |
| b    | eceb   | {e:2,c:1,b:1} | 3         | ❌ shrink |
|      | ceb    | {c:1,e:1,b:1} | 3         | ❌ shrink |
|      | eb     | {e:1,b:1}     | 2         | ✅ valid  |
*/

//
/*
	Complexity
Time: O(n)

Space: O(k)
*/

/*
	replacements_needed = window_size - maxFreq
	We are NOT actually replacing characters in code
👉 We are just checking if replacement is possible

| Right | Window | maxFreq | Window Size | Replace Needed | Valid?   |
| ----- | ------ | ------- | ----------- | -------------- | -------- |
| 0     | A      | 1       | 1           | 0              | ✅        |
| 1     | AA     | 2       | 2           | 0              | ✅        |
| 2     | AAB    | 2       | 3           | 1              | ✅        |
| 3     | AABA   | 3       | 4           | 1              | ✅        |
| 4     | AABAB  | 3       | 5           | 2              | ❌ shrink |
| 5     | ABABB  | 3       | 5           | 2              | ❌ shrink |
| 6     | BABBA  | 3       | 5           | 2              | ❌ shrink |


Time: O(n)

Space: O(1) (26 uppercase letters)

*/

package main

import "fmt"

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)

	left := 0
	maxFreq := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {

		// Step 1: Add current character
		count[s[right]]++

		// Step 2: Track max frequency
		if count[s[right]] > maxFreq {
			maxFreq = count[s[right]]
		}

		// Step 3: Check if window invalid
		for (right-left+1)-maxFreq > k {
			count[s[left]]--
			left++
		}

		// Step 4: Update result
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}


/*

	Minimum Window Substring (HARD

	Find smallest substring of s 
	that contains ALL characters of t 
	with required frequency

	Time: O(n)

	Space: O(k) (chars in t)

	s = "AAABBC"
	t = "AABC"


	A → 2 times
	B → 1
	C → 1

	*/



func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := 0
	needCount := len(need)

	window := make(map[byte]int)

	left := 0
	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		if val, ok := need[c]; ok && window[c] == val {
			have++
		}

		// shrink window
		for have == needCount {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			window[s[left]]--

			if val, ok := need[s[left]]; ok && window[s[left]] < val {
				have--
			}

			left++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}

func main() {
	fmt.Println(characterReplacement("AABABBA", 1)) // 4
	fmt.Println(characterReplacement("ABAB", 2))    // 4
	fmt.Println(characterReplacement("AAAA", 2))    // 4
	fmt.Println(characterReplacement("ABCDE", 1))   // 2
}
