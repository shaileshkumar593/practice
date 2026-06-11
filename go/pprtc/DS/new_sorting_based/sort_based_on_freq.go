package main

import (
	"fmt"
	"sort"
)

func sortByFrequency(arr []int) []int {
	// Step 1: Count frequency of each element
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	fmt.Println("req : ", freq)

	// Step 2: Copy unique elements to a slice for sorting
	unique := make([]int, 0, len(freq))
	for num := range freq {
		fmt.Println(num)
		unique = append(unique, num)
	}

	// Step 3: Sort unique elements by frequency descending, then by value ascending
	sort.Slice(unique, func(i, j int) bool {
		if freq[unique[i]] == freq[unique[j]] {
			return unique[i] < unique[j] // smaller element first if same frequency
		}
		return freq[unique[i]] > freq[unique[j]] // higher frequency first
	})

	// Step 4: Build result array based on frequency
	result := make([]int, 0, len(arr))
	for _, num := range unique {
		for i := 0; i < freq[num]; i++ {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	arr := []int{4, 5, 6, 5, 4, 3, 4, 2, 2, 2, 6}
	fmt.Println("Original array:", arr)

	sorted := sortByFrequency(arr)
	fmt.Println("Sorted by frequency:", sorted)
}

/*

Important Map Iteration Variants
1️⃣ Keys only
for key := range freq {
}
2️⃣ Key and value
for key, value := range freq {
	fmt.Println(key, value)
}
3️⃣ Ignore key
for _, value := range freq {
}

*/
