package main

// Find the median of two sorted arrays.
import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	n := len(merged)
	if n%2 == 0 {
		return float64(merged[n/2-1]+merged[n/2]) / 2 // even
	}
	return float64(merged[n/2]) // odd
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // 2
}

/*
	| Dataset length | Median formula                  |
| -------------- | ------------------------------- |
| Odd ((n))      | (x_{(n+1)/2})                   |
| Even ((n))     | (\frac{x_{n/2} + x_{n/2+1}}{2}) |

	Complexity:

		Time: O((m+n) log(m+n))
		Space: O(m+n)



*/
