package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)

	// mapping closing -> opening
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {

		// If closing bracket
		if open, exists := pairs[ch]; exists {

			// Stack empty OR mismatch
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}

			// Pop
			stack = stack[:len(stack)-1]

		} else {
			// Opening bracket
			stack = append(stack, ch)
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
