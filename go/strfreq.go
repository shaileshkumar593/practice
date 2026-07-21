package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Value string
	Freq  int
	First int
}

func rearrange(arr []string) []string {
	freq := make(map[string]int)
	first := make(map[string]int)

	// Count frequency and record first occurrence
	for i, s := range arr {
		freq[s]++
		if _, ok := first[s]; !ok {
			first[s] = i
		}
	}

	// Create unique items
	items := make([]Item, 0, len(freq))
	for k, v := range freq {
		items = append(items, Item{
			Value: k,
			Freq:  v,
			First: first[k],
		})
	}

	// Sort by frequency descending,
	// then by first occurrence ascending
	sort.Slice(items, func(i, j int) bool {
		if items[i].Freq == items[j].Freq {
			return items[i].First < items[j].First
		}
		return items[i].Freq > items[j].Freq
	})

	// Build result
	result := make([]string, 0, len(arr))
	for _, item := range items {
		for i := 0; i < item.Freq; i++ {
			result = append(result, item.Value)
		}
	}

	return result
}

func main() {
	input := []string{
		"apple",
		"banana",
		"apple",
		"orange",
		"banana",
		"apple",
		"grape",
	}

	fmt.Println(rearrange(input))
}