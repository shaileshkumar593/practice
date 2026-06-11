/*

What is Bucket Sort?
Bucket Sort is a distribution sort:

👉 You divide elements into buckets,
👉 sort each bucket,
👉 then merge.

🔹 When to use?
Numbers are uniformly distributed

Range is known

You want O(n) average time


 Idea:
Create n buckets for n elements

Map element → bucket index

Sort inside each bucket


*/


package main

import (
	"fmt"
	"sort"
)

func bucketSortFloat(arr []float64) []float64 {
	n := len(arr)

	buckets := make([][]float64, n)

	// distribute
	for _, num := range arr {
		index := int(num * float64(n))
		buckets[index] = append(buckets[index], num)
	}

	// sort each bucket
	for i := 0; i < n; i++ {
		sort.Float64s(buckets[i])
	}

	// merge
	result := []float64{}
	for i := 0; i < n; i++ {
		result = append(result, buckets[i]...)
	}

	return result
}

func main() {
	arr := []float64{0.78, 0.12, 0.45, 0.91, 0.33}
	fmt.Println(bucketSortFloat(arr))
}

/*
Range = (max - min) / number_of_buckets

Index = (value - min) / range

For floats (0–1) → value * n

Choose buckets carefully for performance
*/
