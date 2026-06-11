package practice

import (
	"math"
	"sort"
	"strings"
)

// 1. First Unique Character in a String
// Return index of first non-repeating char or -1 if none.
// Time: O(n), Space: O(1) (alphabet fixed)
func FirstUniqueChar(s string) int {
	count := [256]int{} // all possible characters array
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if count[s[i]] == 1 {
			return i
		}
	}
	return -1
}

// 2. Check if Two Strings Are Anagrams
// Time: O(n), Space: O(1) (alphabet fixed)
// same characters with the same frequencies, but possibly in a different order.
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [256]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

// 3. Ransom Note
// Return true if ransomNote can be constructed from magazine.
// Time: O(n+m), Space: O(1)
func CanConstruct(ransomNote, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := [256]int{}
	for i := 0; i < len(magazine); i++ {
		cnt[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		cnt[ransomNote[i]]--
		if cnt[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}

// 4. Count the Number of Vowels
// Time: O(n), Space: O(1)
func CountVowels(s string) int {
	v := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	count := 0
	for _, r := range s {
		if v[r] {
			count++
		}
	}
	return count
}

// 5. Reverse a String (in-place)
// Accepts []byte (slice of characters) and reverses it.
func ReverseInPlace(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

// 6. Find All Extra Characters (t is s plus multiple added chars, shuffled)
// We return a string composed of characters that were added (in the order they appear in t).
// Time: O(n+m), Space: O(1)
func FindAllExtraCharacters(s, t string) string {
	cnt := [256]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}
	var sb strings.Builder
	for i := 0; i < len(t); i++ {
		if cnt[t[i]] > 0 {
			cnt[t[i]]--
		} else {
			sb.WriteByte(t[i])
		}
	}
	return sb.String()
}

// 7. Missing Number (XOR)
// nums contains n numbers in range 0..n with one missing.
func MissingNumber(nums []int) int {
	n := len(nums)
	x := 0
	for i := 0; i < n; i++ {
		x ^= nums[i]
	}
	for i := 0; i <= n; i++ {
		x ^= i
	}
	return x
}

// 8. Valid Palindrome With Deletion of at most one char
// Time: O(n), Space: O(1)
func ValidPalindromeWithOneDeletion(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
			continue
		}
		// try skip left or skip right
		if isPalRange(s, i+1, j) || isPalRange(s, i, j-1) {
			return true
		}
		return false
	}
	return true
}
func isPalRange(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 9. Longest Substring Without Repeating Characters
// Sliding window using last index map
// Time: O(n), Space: O(1)
func LengthOfLongestUniqueSubstring(s string) int {
	last := [256]int{}
	for i := range last {
		last[i] = -1
	}
	start := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if last[c] >= start {
			start = last[c] + 1
		}
		last[c] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}

// 10. Character Frequency Sort (most frequent first). If equal frequency, sort by rune value.
// Time: O(n + k log k) where k = distinct chars, Space: O(k)
func FrequencySort(s string) string {
	if s == "" {
		return ""
	}
	cnt := map[rune]int{}
	for _, r := range s {
		cnt[r]++
	}
	type kv struct {
		r rune
		c int
	}
	arr := make([]kv, 0, len(cnt))
	for r, c := range cnt {
		arr = append(arr, kv{r, c})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].c == arr[j].c {
			return arr[i].r < arr[j].r
		}
		return arr[i].c > arr[j].c
	})
	var sb strings.Builder
	for _, pair := range arr {
		for k := 0; k < pair.c; k++ {
			sb.WriteRune(pair.r)
		}
	}
	return sb.String()
}

// 11. Group Anagrams
// Time: O(n * L log L) if sorting each string length L, Space: O(nL)
func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	m := map[string][]string{}
	for _, s := range strs {
		// sort runes for stable key
		r := []rune(s)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		key := string(r)
		m[key] = append(m[key], s)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 12. Minimum Window Substring
// Standard sliding-window with counts.
// Time: O(n + m), Space: O(1) for alphabet limited
func MinWindow(s, t string) string {
	if len(t) == 0 {
		return ""
	}
	required := [256]int{}
	for i := 0; i < len(t); i++ {
		required[t[i]]++
	}
	formed := 0
	windowCounts := [256]int{}
	requiredUnique := 0
	for i := 0; i < 256; i++ {
		if required[i] > 0 {
			requiredUnique++
		}
	}
	l, r := 0, 0
	minLen := math.MaxInt32
	minL := 0
	for r < len(s) {
		c := s[r]
		windowCounts[c]++
		if windowCounts[c] == required[c] {
			formed++
		}
		for l <= r && formed == requiredUnique {
			if r-l+1 < minLen {
				minLen = r - l + 1
				minL = l
			}
			// remove s[l]
			cl := s[l]
			windowCounts[cl]--
			if windowCounts[cl] < required[cl] {
				formed--
			}
			l++
		}
		r++
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minL : minL+minLen]
}

// 13. Check if a String Is a Permutation of a Palindrome
// Ignore spaces and case. At most one character may have an odd count.
// Time: O(n), Space: O(1)
func IsPermutationOfPalindrome(s string) bool {
	cnt := map[rune]int{}
	for _, r := range s {
		if r == ' ' {
			continue
		}
		r = rune(strings.ToLower(string(r))[0]) // normalize; simpler than import unicode
		cnt[r]++
	}
	odd := 0
	for _, v := range cnt {
		if v%2 != 0 {
			odd++
			if odd > 1 {
				return false
			}
		}
	}
	return true
}

// 14. Detect Duplicate Characters Without Extra Space (only lowercase a-z)
// Use bit vector. Return true if duplicate exists.
// Time: O(n), Space: O(1)
func HasDuplicateLowercaseNoExtraSpace(s string) bool {
	var bitset uint32 = 0
	for _, r := range s {
		if r < 'a' || r > 'z' {
			// for this version we assume only a-z; if other chars appear, treat them as not allowed
			continue
		}
		pos := r - 'a'
		mask := uint32(1) << pos
		if bitset&mask != 0 {
			return true
		}
		bitset |= mask
	}
	return false
}

// 15. Find the Extra Character in Multiple Strings
// Given array where all strings are identical except one has one extra char.
// XOR every byte across all strings -> leftover is extra char.
// Time: O(total chars), Space: O(1)
func FindExtraCharAmongMultipleStrings(arr []string) byte {
	var x byte = 0
	for _, s := range arr {
		for i := 0; i < len(s); i++ {
			x ^= s[i]
		}
	}
	return x
}
