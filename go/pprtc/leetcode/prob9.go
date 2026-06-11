package main

import (
	"fmt"
	"math"
)

/*
	Problem:
		Convert string to 32-bit signed integer.

*/

func myAtoi(s string) int {
	i, sign, n := 0, 1, len(s)
	for i < n && s[i] == ' ' {
		i++
	} // skip spaces

	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	res := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		res = res*10 + digit
		i++
	}
	return res * sign
}

func main() {
	fmt.Println(myAtoi("42"))      // 42
	fmt.Println(myAtoi("   -42"))  // -42
	fmt.Println(myAtoi("4193abc")) // 4193
}

/*

	Explanation:

Skip whitespace, handle optional sign.

Read digits until non-digit.

Handle overflow for 32-bit signed integer.

Complexity:

Time: O(n)

Space: O(1)
	Complexity:

		Time: O(n)

		Space: O(1)

*/
