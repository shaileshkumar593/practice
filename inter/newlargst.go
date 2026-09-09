package main

func picTwo(a []int) {
	var firstlargest, secondlargest int

	for _, val := range a {
		if val > firstlargest {
			firstlargest = val
		} else if firstlargest > val && val > secondlargest {
			secondlargest = val
		}
	}
}

func main() {
	s := []int{7, 2, 5, 6, 3}

	// s := []int{5,10,0}  issue is first = 10 second = 0
	picTwo(s)
}

// what is issue in code
