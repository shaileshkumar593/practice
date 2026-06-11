package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
	1. Basic String Operations
		Find length of a string.

		Convert string to uppercase/lowercase.

		Compare two strings.

		Concatenate strings.

		Trim spaces, prefix, or suffix.

		Check prefix or suffix of a string.

		Split string into words.

		Join array/slice of strings.

2. Character-based Operations
		Count vowels and consonants.

		Count digits and special characters.

		Detect if a string contains only letters or digits.

		Check for palindrome.

		Reverse a string.

		Swap two characters in a string.

		Convert string to []rune for Unicode-safe manipulation.

3. Substring & Pattern Matching
		Generate all substrings.

		Generate substrings without repeated characters.

		Find first occurrence of a substring.

		Check if a substring exists.

		Count occurrences of a substring.

		Replace substring with another string.

		Find longest common prefix.

		Find longest repeated substring.

4. Unicode & Rune Operations
		Iterate over string rune by rune.

		Decode runes with utf8.DecodeRuneInString.

		Encode/decode string to/from UTF-8 bytes.

		Count runes in a string.

		Detect if string contains only ASCII characters.

		Swap runes in a string.

5. Performance-focused Operations
		Build string with strings.Builder.

		Use []byte or []rune for in-place modification.

		Pre-allocate slice for large string operations.

		Join multiple strings efficiently with strings.Join.

		Copy slices with copy function.

6. Searching & Sorting
		Find index of a character or substring.

		Check if string contains only a set of allowed characters.

		Sort characters in a string.

		Remove duplicate characters from string.

		Count frequency of each character.

7. Encoding & Conversion
		Convert string to integer/float.

		Convert integer/float to string.

		Convert string to byte slice and vice versa.

		Encode string to Base64.

		Decode Base64 to string.

		Escape special characters.

8. Misc / Advanced
		Check anagrams.

		Find permutations of a string.

		Check if string has all unique characters.

		Detect palindromic substrings.

		Count words in a string.

		Remove spaces or extra whitespace.

		Count sentences or paragraphs.

		Convert string to title case or camelCase.

		Tokenize string using delimiters.

		Implement simple string compression (like aaabbc -> a3b2c1).


*/

func main() {
	fmt.Println("=== STRING OPERATIONS DEMO ===")

	s1 := "Hello World 123!"
	s2 := "GoLang"

	fmt.Println("\n--- 1. Basic String Operations ---")
	fmt.Println("Length:", len(s1))
	fmt.Println("Uppercase:", strings.ToUpper(s1))
	fmt.Println("Lowercase:", strings.ToLower(s1))
	fmt.Println("Compare s1 & s2:", strings.Compare(s1, s2))
	fmt.Println("Concatenate:", s1+s2)
	fmt.Println("Trim spaces:", strings.TrimSpace("  Hello  "))
	fmt.Println("Trim prefix:", strings.TrimPrefix("GoLang", "Go"))
	fmt.Println("Trim suffix:", strings.TrimSuffix("GoLang", "Lang"))
	fmt.Println("HasPrefix:", strings.HasPrefix(s1, "Hello"))
	fmt.Println("HasSuffix:", strings.HasSuffix(s1, "!"))
	fmt.Println("Split into words:", strings.Fields(s1))
	fmt.Println("Join words:", strings.Join([]string{"Go", "is", "fun"}, " "))

	fmt.Println("\n--- 2. Character-based Operations ---")
	vowels, consonants := 0, 0
	digits, special := 0, 0
	for _, r := range s1 {
		if unicode.IsLetter(r) {
			switch unicode.ToLower(r) {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			default:
				consonants++
			}
		} else if unicode.IsDigit(r) {
			digits++
		} else {
			special++
		}
	}
	fmt.Println("Vowels:", vowels, "Consonants:", consonants, "Digits:", digits, "Special:", special)

	fmt.Println("Palindrome 'level':", isPalindrome("level"))
	fmt.Println("Reverse s1:", reverseString(s1))
	fmt.Println("Swap first two chars of s2:", swapChars(s2, 0, 1))

	fmt.Println("\n--- 3. Substring & Pattern Matching ---")
	fmt.Println("All substrings of 'abc':", allSubstrings("abc"))
	fmt.Println("Substring 'World' exists:", strings.Contains(s1, "World"))
	fmt.Println("Index of 'World':", strings.Index(s1, "World"))
	fmt.Println("Count 'l' in s1:", strings.Count(s1, "l"))
	fmt.Println("Replace 'Hello' with 'Hi':", strings.ReplaceAll(s1, "Hello", "Hi"))
	fmt.Println("Longest common prefix of ['flower','flow','flight']:", longestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Println("\n--- 4. Unicode & Rune Operations ---")
	fmt.Println("Rune count in s1:", utf8.RuneCountInString(s1))
	for i, r := range s1 {
		fmt.Printf("Rune %d: %c\n", i, r)
	}
	fmt.Println("Is ASCII:", isASCII(s1))

	fmt.Println("\n--- 5. Performance-focused Operations ---")
	var builder strings.Builder
	builder.WriteString("Go")
	builder.WriteString(" Builder")
	fmt.Println("Using strings.Builder:", builder.String())

	bytesSlice := []byte("Byte slice example")
	copySlice := make([]byte, len(bytesSlice))
	copy(copySlice, bytesSlice)
	fmt.Println("Copied byte slice:", string(copySlice))

	fmt.Println("\n--- 6. Searching & Sorting ---")
	fmt.Println("Sort characters in 'dcba':", sortString("dcba"))
	fmt.Println("Remove duplicate chars in 'hello':", removeDuplicateChars("hello"))
	fmt.Println("Frequency count of chars in 'hello':", charFrequency("hello"))

	fmt.Println("\n--- 7. Encoding & Conversion ---")
	numStr := "100"
	num, _ := strconv.Atoi(numStr)
	fmt.Println("String to int:", num)
	fmt.Println("Int to string:", strconv.Itoa(num))
	floatStr := "12.34"
	f, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Println("String to float:", f)
	fmt.Println("Float to string:", strconv.FormatFloat(f, 'f', 2, 64))
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(s1))
	fmt.Println("Base64 encoded:", base64Encoded)
	decoded, _ := base64.StdEncoding.DecodeString(base64Encoded)
	fmt.Println("Base64 decoded:", string(decoded))

	fmt.Println("\n--- 8. Misc / Advanced ---")
	fmt.Println("Check anagram 'listen' & 'silent':", isAnagram("listen", "silent"))
	fmt.Println("Check unique chars in 'abc':", hasUniqueChars("abc"))
	fmt.Println("Count words in s1:", len(strings.Fields(s1)))
	fmt.Println("Remove extra spaces:", strings.Join(strings.Fields("  Go   is   fun "), " "))
	fmt.Println("Title case:", strings.Title("go is fun"))
	fmt.Println("Simple compression 'aaabbc':", stringCompression("aaabbc"))
}

// Reverse string using runes
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Swap two characters in string using runes
func swapChars(s string, i, j int) string {
	r := []rune(s)
	if i < len(r) && j < len(r) {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Check palindrome
func isPalindrome(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

// Generate all substrings
func allSubstrings(s string) []string {
	var res []string
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		for j := i + 1; j <= len(r); j++ {
			res = append(res, string(r[i:j]))
		}
	}
	return res
}

// Longest common prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs[1:] {
		for !strings.HasPrefix(s, prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// Check if string is ASCII only
func isASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

// Sort characters in string
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

// Remove duplicate characters
func removeDuplicateChars(s string) string {
	r := []rune(s)
	result := []rune{}
	for i, ch := range r {
		unique := true
		for j := 0; j < i; j++ {
			if r[j] == ch {
				unique = false
				break
			}
		}
		if unique {
			result = append(result, ch)
		}
	}
	return string(result)
}

// Count frequency of characters
func charFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	return freq
}

// Check anagram
func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[rune]int)
	for _, r := range a {
		count[r]++
	}
	for _, r := range b {
		count[r]--
		if count[r] < 0 {
			return false
		}
	}
	return true
}

// Check if all characters are unique
func hasUniqueChars(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] == r[j] {
				return false
			}
		}
	}
	return true
}

// Simple string compression
func stringCompression(s string) string {
	r := []rune(s)
	var b bytes.Buffer
	count := 1
	for i := 0; i < len(r); i++ {
		if i+1 < len(r) && r[i] == r[i+1] {
			count++
		} else {
			b.WriteRune(r[i])
			if count > 1 {
				b.WriteString(strconv.Itoa(count))
			}
			count = 1
		}
	}
	return b.String()
}
