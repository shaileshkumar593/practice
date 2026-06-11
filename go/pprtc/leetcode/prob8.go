package main

/*
	Problem:
		Reverse digits of an integer, handling overflow.
*/

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	rev := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && digit < -8) {
			return 0
		}
		rev = rev*10 + digit
	}
	return rev
}

func main() {
	fmt.Println(reverse(123))  // 321
	fmt.Println(reverse(-123)) // -321
}

/*

	Step-by-Step Explanation

	func reverse(x int) int {
		rev := 0
	rev is the variable that will hold the reversed number.

		for x != 0 {
	Loop until all digits are processed.

	Each iteration extracts the last digit of x and appends it to rev.

			digit := x % 10
			x /= 10
	% 10 gets the last digit of x.

	/= 10 removes the last digit from x.

	Example: x=123 → digit=3, x=12

			if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && digit > 7) {
				return 0
			}
			if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && digit < -8) {
				return 0
			}
	Overflow check:

	math.MaxInt32 = 2147483647

	math.MinInt32 = -2147483648

	When appending a new digit: rev = rev*10 + digit, we must ensure this does not overflow 32-bit integer limits.

	rev > MaxInt32/10 → multiplying by 10 will overflow

	rev == MaxInt32/10 && digit > 7 → adding the digit will overflow

	Similarly for negative numbers with MinInt32

	If overflow would occur, the function returns 0.

			rev = rev*10 + digit
	Append the digit to rev.

	Example: rev=12, digit=3 → rev=123

		}
		return rev
	}
	After the loop ends, return the reversed integer.

*/
