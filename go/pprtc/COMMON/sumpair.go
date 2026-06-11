package main

import "fmt"

func main() {
	//arr := []int{5, 1, 7, -1, 5}
	arr := []int{5, 1, 4, 3, 6, 2, 0, 3}
	//arr := []int{-1, -2, 0, 4, 3}  sum =0
	sum := 6
	count1 := 0
	mapping := make(map[int]int, len(arr))

	for i, val := range arr {
		if mapping[sum-val] == 0 {
			if i == 0 {
				mapping[val] = 99999
			} else {
				mapping[val] = i
			}

			fmt.Println(mapping)
		} else {
			if mapping[sum-val] == 99999 {
				fmt.Println("Found Pair ", arr[mapping[0]], "---->", arr[i])
			} else {
				fmt.Println("Found Pair ", arr[mapping[sum-val]], "---->", arr[i])
			}

			count1 = count1 + 1
		}
	}

	if count1 == 0 {
		fmt.Println("Pair Not Exist")
	} else {
		fmt.Println("Pair  Exist")
	}

}
