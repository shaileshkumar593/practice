package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	// mapping closing → opening
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {

		// if opening → push
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)

		} else {
			// if closing → check stack

			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if top != pairs[ch] {
				return false
			}

			// pop
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
