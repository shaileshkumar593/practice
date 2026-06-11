package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "₹50"

	fmt.Println("---- Unicode Checks ----")
	for _, r := range s {
		fmt.Printf("%c: Letter=%v Digit=%v Space=%v Upper=%v Lower=%v\n",
			r,
			unicode.IsLetter(r),
			unicode.IsDigit(r),
			unicode.IsSpace(r),
			unicode.IsUpper(r),
			unicode.IsLower(r),
		)
	}

	fmt.Println("\n---- UTF8 Functions ----")
	fmt.Println("Bytes length:", len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	r, size := utf8.DecodeRuneInString(s)
	fmt.Printf("First rune: %c, bytes used: %d\n", r, size)
	fmt.Println("Rune length:", utf8.RuneLen(r))
	fmt.Println("Valid UTF8:", utf8.ValidString(s))

	fmt.Println("\n---- String Conversion ----")
	fmt.Println("ToUpper:", strings.ToUpper("élève"))
	fmt.Println("ToLower:", strings.ToLower("ÉLÈVE"))
}

/*
	| Function                | Package   | Description                  | Example                          |
| ----------------------- | --------- | ---------------------------- | -------------------------------- |
| `IsLetter(r)`           | `unicode` | Checks if rune is a letter   | `unicode.IsLetter('₹')`          |
| `IsDigit(r)`            | `unicode` | Checks if rune is digit      | `unicode.IsDigit('9')`           |
| `IsSpace(r)`            | `unicode` | Checks if rune is whitespace | `unicode.IsSpace('\t')`          |
| `ToUpper(r)`            | `unicode` | Convert rune to uppercase    | `unicode.ToUpper('a')`           |
| `DecodeRuneInString(s)` | `utf8`    | Decode first rune            | `utf8.DecodeRuneInString("₹50")` |
| `RuneCountInString(s)`  | `utf8`    | Count runes in string        | `utf8.RuneCountInString("₹50")`  |
| `RuneLen(r)`            | `utf8`    | Bytes per rune               | `utf8.RuneLen('₹')`              |
| `EncodeRune(p, r)`      | `utf8`    | Encode rune to bytes         | `utf8.EncodeRune(buf, '₹')`      |
| `ValidString(s)`        | `utf8`    | Check UTF-8 validity         | `utf8.ValidString("₹50")`        |
*/

/*

	Packages Involved
Most rune-related functions live in:

unicode

unicode/utf8

strings (some helper functions use runes internally)

🧩 1. unicode Package — Rune Classification and Conversion
✅ 1.1 unicode.IsLetter(r rune) bool
Returns whether the rune is a letter (A–Z, a–z, or any Unicode letter).

fmt.Println(unicode.IsLetter('A'))   // true
fmt.Println(unicode.IsLetter('₹'))   // true
fmt.Println(unicode.IsLetter('9'))   // false
✅ 1.2 unicode.IsDigit(r rune) bool
Checks if the rune is a digit (0–9, or other Unicode digits).

fmt.Println(unicode.IsDigit('5'))   // true
fmt.Println(unicode.IsDigit('²'))   // true
fmt.Println(unicode.IsDigit('a'))   // false
✅ 1.3 unicode.IsSpace(r rune) bool
Checks for space-like characters (space, tab, newline, etc).

fmt.Println(unicode.IsSpace(' '))   // true
fmt.Println(unicode.IsSpace('\t'))  // true
fmt.Println(unicode.IsSpace('A'))   // false
✅ 1.4 unicode.IsUpper(r rune) bool and unicode.IsLower(r rune) bool
fmt.Println(unicode.IsUpper('A')) // true
fmt.Println(unicode.IsLower('a')) // true
✅ 1.5 unicode.ToUpper(r rune), unicode.ToLower(r rune), unicode.ToTitle(r rune)
Convert a rune’s case according to Unicode rules.

fmt.Printf("%c\n", unicode.ToUpper('a')) // A
fmt.Printf("%c\n", unicode.ToLower('A')) // a
fmt.Printf("%c\n", unicode.ToTitle('ß')) // ẞ
✅ 1.6 unicode.In(r rune, ranges ...*RangeTable) bool
Check if a rune belongs to a Unicode category like unicode.Latin, unicode.Han, etc.

fmt.Println(unicode.In('₹', unicode.Devanagari)) // true
fmt.Println(unicode.In('中', unicode.Han))       // true
🧩 2. unicode/utf8 Package — Encoding & Decoding
These deal with UTF-8 byte <-> rune conversion.

✅ 2.1 utf8.RuneCountInString(s string) int
Counts the number of runes (characters) in a UTF-8 string.

s := "₹50"
fmt.Println(len(s))                     // 5 (bytes)
fmt.Println(utf8.RuneCountInString(s))  // 3 (runes)
✅ 2.2 utf8.DecodeRuneInString(s string) (r rune, size int)
Decodes the first rune and returns how many bytes it took.

r, size := utf8.DecodeRuneInString("₹50")
fmt.Printf("%c %U %d\n", r, r, size)
Output:

₹ U+20B9 3
✅ 2.3 utf8.EncodeRune(p []byte, r rune) int
Encodes a rune into a UTF-8 byte slice.

buf := make([]byte, 3)
n := utf8.EncodeRune(buf, '₹')
fmt.Println(buf[:n]) // [226 130 185]
✅ 2.4 utf8.Valid([]byte) and utf8.ValidString(string)
Checks whether the byte sequence or string is valid UTF-8.

fmt.Println(utf8.ValidString("₹50")) // true
fmt.Println(utf8.Valid([]byte{0xff})) // false
✅ 2.5 utf8.RuneLen(r rune) int
Returns how many bytes a rune takes in UTF-8.

fmt.Println(utf8.RuneLen('A')) // 1
fmt.Println(utf8.RuneLen('₹')) // 3
fmt.Println(utf8.RuneLen('𐍈')) // 4
🧩 3. strings Package — Rune-Aware Operations
Some high-level functions handle runes automatically.

✅ 3.1 strings.ToUpper() / strings.ToLower()
Convert entire strings with Unicode support.

fmt.Println(strings.ToUpper("élève")) // ÉLÈVE
✅ 3.2 strings.Runes(s string) []rune
Convert string to a rune slice.

runes := []rune("₹50")
fmt.Println(runes) // [8377 53 48]

*/
