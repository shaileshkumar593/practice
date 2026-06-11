package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	// Input string
	s := "GoLang ₹50"
	fmt.Println("Original string:", s)
	fmt.Println("Length in bytes:", len(s))
	fmt.Println("Length in runes:", utf8.RuneCountInString(s))

	// ------------- Basic checks ----------------
	fmt.Println("Contains 'Go':", strings.Contains(s, "Go"))
	fmt.Println("Count of 'g':", strings.Count(strings.ToLower(s), "g"))
	fmt.Println("HasPrefix 'Go':", strings.HasPrefix(s, "Go"))
	fmt.Println("HasSuffix '50':", strings.HasSuffix(s, "50"))
	fmt.Println("Index of 'Lang':", strings.Index(s, "Lang"))
	fmt.Println("Last Index of 'g':", strings.LastIndex(s, "g"))

	// ------------- Case conversion ----------------
	fmt.Println("ToUpper:", strings.ToUpper(s))
	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("Title:", strings.Title(strings.ToLower(s))) // First char uppercase

	// ------------- Trim and split ----------------
	str2 := "   Hello GoLang   "
	fmt.Println("TrimSpace:", strings.TrimSpace(str2))
	fmt.Println("Trim '-':", strings.Trim("---abc---", "-"))
	fmt.Println("Split:", strings.Split("a,b,c", ","))
	fmt.Println("SplitN:", strings.SplitN("a,b,c,d", ",", 3))

	// ------------- Replace / Repeat ----------------
	fmt.Println("Replace:", strings.Replace("hello world", "l", "L", 2))
	fmt.Println("ReplaceAll:", strings.ReplaceAll("hello world", "l", "L"))
	fmt.Println("Repeat:", strings.Repeat("abc", 3))

	// ------------- Compare / EqualFold ----------------
	fmt.Println("Compare 'abc' vs 'ABC':", strings.Compare("abc", "ABC"))
	fmt.Println("EqualFold 'abc' vs 'ABC':", strings.EqualFold("abc", "ABC"))

	// ------------- Fields / IndexFunc ----------------
	fmt.Println("Fields:", strings.Fields("a b c"))
	fmt.Println("FieldsFunc split by ',' or ';':",
		strings.FieldsFunc("a,b;c", func(r rune) bool { return r == ',' || r == ';' }))

	fmt.Println("IndexFunc digit:", strings.IndexFunc("abc123", func(r rune) bool { return r >= '0' && r <= '9' }))

	// ------------- Strings Builder ----------------
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	builder.WriteRune('!')
	fmt.Println("Builder Result:", builder.String())
	fmt.Println("Builder Length:", builder.Len())
	builder.Reset()
	fmt.Println("Builder after Reset:", builder.String())

	// ------------- Rune iteration ----------------
	fmt.Println("Iterate string by rune:")
	for i, r := range s {
		fmt.Printf("Index %d: %c\n", i, r)
	}

	// ------------- Conversion with strconv ----------------
	numStr := "123"
	num, _ := strconv.Atoi(numStr)
	fmt.Println("String to int:", num, " + 10 =", num+10)
	str := strconv.Itoa(456)
	fmt.Println("Int to string:", str)

	// ------------- Generate substring without repetition example -------------
	fmt.Println("Substrings without repeated chars from 'abc':")
	generateSubstrings("abc")
}

// ---------------- Helper function ----------------
func generateSubstrings(s string) {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if !isUnique(runes, i, j) {
				break
			}
			fmt.Println(string(runes[i : j+1]))
		}
	}
}

func isUnique(runes []rune, start, end int) bool {
	for i := start; i < end; i++ {
		if runes[i] == runes[end] {
			return false
		}
	}
	return true
}

/*

------------------- STRING FUNCTIONS SUMMARY TABLE -------------------

| Function / Method           | Package    | Description                                           |
|-----------------------------|------------|-------------------------------------------------------|
| len(s)                       | builtin    | Length in bytes                                      |
| utf8.RuneCountInString(s)    | utf8       | Length in runes                                      |
| strings.Contains             | strings    | Check if substring exists                            |
| strings.Count                | strings    | Count occurrences                                    |
| strings.HasPrefix / HasSuffix| strings    | Check prefix/suffix                                  |
| strings.Index / LastIndex    | strings    | Find index                                           |
| strings.ToUpper / ToLower    | strings    | Case conversion                                      |
| strings.Title                | strings    | Capitalize first letter                              |
| strings.Trim / TrimSpace     | strings    | Remove unwanted characters or spaces                 |
| strings.Split / SplitN       | strings    | Split string into slice                               |
| strings.Replace / ReplaceAll | strings    | Replace substring                                    |
| strings.Repeat               | strings    | Repeat string                                        |
| strings.Compare              | strings    | Lexicographic comparison                              |
| strings.EqualFold            | strings    | Case-insensitive comparison                           |
| strings.Fields / FieldsFunc  | strings    | Split by spaces or custom function                   |
| strings.IndexFunc / LastIndexFunc | strings | Find index using function                             |
| strings.Builder              | strings    | Efficient string concatenation                        |
| strconv.Atoi / Itoa          | strconv    | Convert string ↔ int                                  |

*/
