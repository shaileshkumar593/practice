package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*

	func Aggregation(m map[int]int, freMin, freMax int) (sum, count, minVal, maxVal int) {
	minVal = math.MaxInt
	maxVal = math.MinInt

	for key, freq := range m {

		// Fast skip if outside range
	// 	if freq < freMin || freq > freMax {
	// 		continue
	// 	}

	// 	sum += key
	// 	count++

	// 	// Update min
	// 	if key < minVal {
	// 		minVal = key
	// 	}
	// 	// Update max
	// 	if key > maxVal {
	// 		maxVal = key
	// 	}
	// }

	// No values found
	// if count == 0 {
	// 	return 0, 0, -1, -1
	// }

	return
}
*/

func Aggregation(m map[int]int, freqmin, freqmax int) (sum int, count int, minval int, maxval int) {

	minVal := math.MaxInt
	maxVal := math.MinInt

	for key, val := range m {

		if val >= freqmin && val <= freqmax {

			sum += key
			count++

			if key < minVal {
				minVal = key
			}
			if key > maxVal {
				maxVal = key
			}

		}

	}

	return sum, count, min, max
}

func main() {

	l := make([]int, 0)

	for i := 0; i < 1000000; i++ {
		l = append(l, rand.Int())
	}

	m := make(map[int]int, 0)

	for _, val := range l {
		if val1, exist := m[val]; exist {
			m[val] = val1 + 1
		} else {
			m[val] = 1
		}
	}

	sum, count, minval, maxval := Aggregation(m, 7, 400)

	fmt.Println(sum, count, minval, maxval)

}
