package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Go😊"

	fmt.Println("===================================")
	fmt.Println("1. range over string (Recommended)")
	fmt.Println("===================================")

	for i, ch := range s {
		fmt.Printf("index=%d char=%c type=%T\n", i, ch, ch)
	}

	fmt.Println("\n==============================")
	fmt.Println("2. Traditional for (bytes)")
	fmt.Println("==============================")

	for i := 0; i < len(s); i++ {
		fmt.Printf("index=%d byte=%d char=%c\n", i, s[i], s[i])
	}

	fmt.Println("\n===================================")
	fmt.Println("3. Convert string to []rune")
	fmt.Println("===================================")

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("index=%d rune=%c type=%T\n",
			i,
			runes[i],
			runes[i],
		)
	}

	fmt.Println("\n===================================")
	fmt.Println("4. range over []rune")
	fmt.Println("===================================")

	for i, ch := range []rune(s) {
		fmt.Printf("index=%d rune=%c\n", i, ch)
	}

	fmt.Println("\n===================================")
	fmt.Println("5. range over []byte")
	fmt.Println("===================================")

	for i, b := range []byte(s) {
		fmt.Printf("index=%d byte=%d\n", i, b)
	}

	fmt.Println("\n===================================")
	fmt.Println("6. Infinite style for")
	fmt.Println("===================================")

	for i := 0; ; i++ {
		if i >= len(s) {
			break
		}

		fmt.Printf("index=%d char=%c\n", i, s[i])
	}

	fmt.Println("\n===================================")
	fmt.Println("7. While loop style")
	fmt.Println("===================================")

	i := 0

	for i < len(s) {
		fmt.Printf("index=%d char=%c\n", i, s[i])
		i++
	}

	fmt.Println("\n===================================")
	fmt.Println("8. UTF8 DecodeRune")
	fmt.Println("===================================")

	temp := s

	for len(temp) > 0 {
		r, size := utf8.DecodeRuneInString(temp)

		fmt.Printf("rune=%c size=%d\n", r, size)

		temp = temp[size:]
	}

	fmt.Println("\n===================================")
	fmt.Println("9. Only character")
	fmt.Println("===================================")

	for _, ch := range s {
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	fmt.Println("\n===================================")
	fmt.Println("10. Only index")
	fmt.Println("===================================")

	for i := range s {
		fmt.Println(i)
	}
}
