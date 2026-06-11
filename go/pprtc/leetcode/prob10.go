package main

/*
	Problem:
		Check if an integer is a palindrome.
*/

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10
}

func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
}

/*
		Explanation:

Negative numbers are never palindrome.

Reverse half of the number and compare with the other half.

Complexity:

Time: O(log10(x))

Space: O(1)

*/
