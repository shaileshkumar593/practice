package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ----------------------------
// String Utilities
// ----------------------------

// Reverse a string safely using runes (Unicode-safe)
func reverseString(s string) string {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r)
}

// Check if string contains only letters and digits
func isAlphaNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func main() {
	s := "Go ₹ 汉 123 hello"

	fmt.Println("Original String:", s)

	// ----------------------------
	// Length & Rune Count
	// ----------------------------
	fmt.Println("\n--- Length ---")
	fmt.Println("Length in bytes:", len(s))                    // raw bytes
	fmt.Println("Length in runes:", utf8.RuneCountInString(s)) // Unicode-safe

	// ----------------------------
	// Iteration
	// ----------------------------
	fmt.Println("\n--- Iteration using runes ---")
	for i, r := range s {
		fmt.Printf("Index %d: %c (U+%04X)\n", i, r, r)
	}

	fmt.Println("\n--- Iteration using utf8.DecodeRuneInString ---")
	str := s
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("Rune: %c, size: %d bytes\n", r, size)
		str = str[size:]
	}

	// ----------------------------
	// Strings Package Functions
	// ----------------------------
	fmt.Println("\n--- strings package functions ---")
	fmt.Println("Contains 'hello':", strings.Contains(s, "hello"))
	fmt.Println("ContainsAny 'Go₹123':", strings.ContainsAny(s, "Go₹123"))
	fmt.Println("HasPrefix 'Go':", strings.HasPrefix(s, "Go"))
	fmt.Println("HasSuffix 'hello':", strings.HasSuffix(s, "hello"))
	fmt.Println("Index of '汉':", strings.Index(s, "汉"))
	fmt.Println("LastIndex of 'l':", strings.LastIndex(s, "l"))
	fmt.Println("Count of 'l':", strings.Count(s, "l"))
	fmt.Println("Split by space:", strings.Split(s, " "))
	fmt.Println("Join slice ['a','b','c'] with '-':", strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println("ToUpper:", strings.ToUpper(s))
	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("TrimSpace '  hello  ':", strings.TrimSpace("  hello  "))
	fmt.Println("Trim '!!!hello!!!' of '!':", strings.Trim("!!!hello!!!", "!"))
	fmt.Println("TrimLeft '!!!hello!!!' of '!':", strings.TrimLeft("!!!hello!!!", "!"))
	fmt.Println("TrimRight '!!!hello!!!' of '!':", strings.TrimRight("!!!hello!!!", "!"))
	fmt.Println("TrimPrefix 'Go Hello' of 'Go':", strings.TrimPrefix("Go Hello", "Go"))
	fmt.Println("TrimSuffix 'Hello Go' of 'Go':", strings.TrimSuffix("Hello Go", "Go"))
	fmt.Println("ReplaceAll 'hello hello' replace 'l'->'x':", strings.ReplaceAll("hello hello", "l", "x"))
	fmt.Println("Repeat 'ha' 3 times:", strings.Repeat("ha", 3))
	fmt.Println("Fields split by whitespace:", strings.Fields("a b c"))
	fmt.Println("SplitAfter 'a,b,c' by ',':", strings.SplitAfter("a,b,c", ","))
	fmt.Println("SplitAfterN 'a,b,c' by ',' limit 2:", strings.SplitAfterN("a,b,c", ",", 2))
	fmt.Println("Compare 'abc' and 'abd':", strings.Compare("abc", "abd")) // -1
	fmt.Println("EqualFold 'Go' and 'go':", strings.EqualFold("Go", "go")) // true

	// ----------------------------
	// strconv Package Functions
	// ----------------------------
	fmt.Println("\n--- strconv package functions ---")
	i, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println("String to Int:", i)
	}

	sInt := strconv.Itoa(456)
	fmt.Println("Int to String:", sInt)

	f, _ := strconv.ParseFloat("12.34", 64)
	fmt.Println("ParseFloat:", f)

	fs := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Println("FormatFloat:", fs)

	b, _ := strconv.ParseBool("true")
	fmt.Println("ParseBool:", b)
	fmt.Println("FormatBool(false):", strconv.FormatBool(false))

	// ----------------------------
	// Rune Operations
	// ----------------------------
	fmt.Println("\n--- Rune Operations ---")
	for _, r := range s {
		fmt.Printf("%c is letter? %t, is digit? %t, is space? %t\n", r, unicode.IsLetter(r), unicode.IsDigit(r), unicode.IsSpace(r))
	}

	// ----------------------------
	// Reverse string
	// ----------------------------
	fmt.Println("\n--- Reverse string ---")
	fmt.Println("Original:", s)
	fmt.Println("Reversed:", reverseString(s))

	// ----------------------------
	// Alphanumeric check
	// ----------------------------
	fmt.Println("\n--- Alphanumeric Check ---")
	alnum := "Go123"
	nonAlnum := "Go ₹ 123"
	fmt.Printf("'%s' is alphanumeric? %t\n", alnum, isAlphaNumeric(alnum))
	fmt.Printf("'%s' is alphanumeric? %t\n", nonAlnum, isAlphaNumeric(nonAlnum))
}

/*

	| Function                                     | Package        | Description                                                   | Example                                                 |
| -------------------------------------------- | -------------- | ------------------------------------------------------------- | ------------------------------------------------------- |
| `len(s)`                                     | builtin        | Returns the length of string in bytes                         | `len("Go₹") -> 4`                                       |
| `utf8.RuneCountInString(s)`                  | `unicode/utf8` | Returns the number of Unicode code points (runes) in string   | `utf8.RuneCountInString("Go₹") -> 3`                    |
| `utf8.DecodeRuneInString(s)`                 | `unicode/utf8` | Decodes first rune in string and returns rune + size in bytes | `r, size := utf8.DecodeRuneInString("Go₹")`             |
| `strings.Contains(s, substr)`                | `strings`      | Checks if `substr` exists in `s`                              | `strings.Contains("hello", "ell") -> true`              |
| `strings.ContainsAny(s, chars)`              | `strings`      | Checks if any char in `chars` exists in `s`                   | `strings.ContainsAny("Go₹123", "₹") -> true`            |
| `strings.HasPrefix(s, prefix)`               | `strings`      | Checks if string starts with `prefix`                         | `strings.HasPrefix("GoLang", "Go") -> true`             |
| `strings.HasSuffix(s, suffix)`               | `strings`      | Checks if string ends with `suffix`                           | `strings.HasSuffix("hello.go", ".go") -> true`          |
| `strings.Index(s, substr)`                   | `strings`      | Returns first index of `substr` or -1                         | `strings.Index("hello", "l") -> 2`                      |
| `strings.LastIndex(s, substr)`               | `strings`      | Returns last index of `substr` or -1                          | `strings.LastIndex("hello", "l") -> 3`                  |
| `strings.Count(s, substr)`                   | `strings`      | Counts occurrences of `substr`                                | `strings.Count("hello", "l") -> 2`                      |
| `strings.Split(s, sep)`                      | `strings`      | Splits string into slice by separator                         | `strings.Split("a,b,c", ",") -> ["a","b","c"]`          |
| `strings.SplitAfter(s, sep)`                 | `strings`      | Splits string but keeps separator                             | `strings.SplitAfter("a,b,c", ",") -> ["a,", "b,", "c"]` |
| `strings.SplitAfterN(s, sep, n)`             | `strings`      | Split string keeping separator, limit to n parts              | `strings.SplitAfterN("a,b,c", ",", 2) -> ["a,", "b,c"]` |
| `strings.Join(slice, sep)`                   | `strings`      | Joins string slice using separator                            | `strings.Join([]string{"a","b"}, "-") -> "a-b"`         |
| `strings.ToUpper(s)`                         | `strings`      | Converts string to uppercase                                  | `strings.ToUpper("go") -> "GO"`                         |
| `strings.ToLower(s)`                         | `strings`      | Converts string to lowercase                                  | `strings.ToLower("GO") -> "go"`                         |
| `strings.TrimSpace(s)`                       | `strings`      | Removes leading & trailing spaces                             | `strings.TrimSpace(" hi ") -> "hi"`                     |
| `strings.Trim(s, cutset)`                    | `strings`      | Removes leading/trailing characters in cutset                 | `strings.Trim("!!!hi!!!", "!") -> "hi"`                 |
| `strings.TrimLeft(s, cutset)`                | `strings`      | Removes leading characters in cutset                          | `strings.TrimLeft("!!!hi", "!") -> "hi"`                |
| `strings.TrimRight(s, cutset)`               | `strings`      | Removes trailing characters in cutset                         | `strings.TrimRight("hi!!!", "!") -> "hi"`               |
| `strings.TrimPrefix(s, prefix)`              | `strings`      | Removes prefix if present                                     | `strings.TrimPrefix("GoLang", "Go") -> "Lang"`          |
| `strings.TrimSuffix(s, suffix)`              | `strings`      | Removes suffix if present                                     | `strings.TrimSuffix("Lang.go", ".go") -> "Lang"`        |
| `strings.ReplaceAll(s, old, new)`            | `strings`      | Replaces all occurrences of old with new                      | `strings.ReplaceAll("hello","l","x") -> "hexxo"`        |
| `strings.Repeat(s, n)`                       | `strings`      | Repeats string n times                                        | `strings.Repeat("ha", 3) -> "hahaha"`                   |
| `strings.Fields(s)`                          | `strings`      | Splits string by whitespace                                   | `strings.Fields("a b c") -> ["a","b","c"]`              |
| `strings.Compare(a, b)`                      | `strings`      | Compares two strings (-1,0,1)                                 | `strings.Compare("a","b") -> -1`                        |
| `strings.EqualFold(a, b)`                    | `strings`      | Case-insensitive comparison                                   | `strings.EqualFold("Go","go") -> true`                  |
| `strconv.Atoi(s)`                            | `strconv`      | Converts string to int                                        | `strconv.Atoi("123") -> 123`                            |
| `strconv.Itoa(i)`                            | `strconv`      | Converts int to string                                        | `strconv.Itoa(456) -> "456"`                            |
| `strconv.ParseFloat(s, bitSize)`             | `strconv`      | Parses float from string                                      | `strconv.ParseFloat("12.34", 64) -> 12.34`              |
| `strconv.FormatFloat(f, fmt, prec, bitSize)` | `strconv`      | Formats float to string                                       | `strconv.FormatFloat(12.34,'f',2,64) -> "12.34"`        |
| `strconv.ParseBool(s)`                       | `strconv`      | Parses boolean from string                                    | `strconv.ParseBool("true") -> true`                     |
| `strconv.FormatBool(b)`                      | `strconv`      | Converts boolean to string                                    | `strconv.FormatBool(false) -> "false"`                  |
| `unicode.IsLetter(r)`                        | `unicode`      | Checks if rune is letter                                      | `unicode.IsLetter('A') -> true`                         |
| `unicode.IsDigit(r)`                         | `unicode`      | Checks if rune is digit                                       | `unicode.IsDigit('5') -> true`                          |
| `unicode.IsSpace(r)`                         | `unicode`      | Checks if rune is space                                       | `unicode.IsSpace(' ') -> true`                          |
*/
