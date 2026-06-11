package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Country struct {
	Name       string
	Population int
}

func calculatePercentages(countries []Country) map[string]float64 {
	total := 0
	for _, c := range countries {
		total += c.Population
	}

	result := make(map[string]float64)
	for _, c := range countries {
		result[c.Name] = (float64(c.Population) / float64(total)) * 100
	}
	return result
}

func pickCountry(countries []Country) string {
	total := 0
	for _, c := range countries {
		total += c.Population
	}

	r := rand.Intn(total)
	fmt.Println("rand val ", r)
	sum := 0

	/*
		r ∈ [0, 4,236,033]
		Think of this as:

		“Pick ONE person randomly from the entire population”

		Each number represents one person’s index.
	*/

	for _, c := range countries {
		sum += c.Population
		if r < sum {
			return c.Name
		}
	}
	return ""
}

func main() {
	rand.Seed(time.Now().UnixNano())

	countries := []Country{
		{"India", 2345678},
		{"USA", 1345678},
		{"UK", 545678},
	}

	// Show calculated percentages
	percentages := calculatePercentages(countries)
	fmt.Println("Percentages:")
	for k, v := range percentages {
		fmt.Printf("%s → %.2f%%\n", k, v)
	}

	// Run weighted selection many times
	count := map[string]int{}
	for i := 0; i < 100000; i++ {
		c := pickCountry(countries)
		count[c]++
	}

	fmt.Println("\nSelection result:")
	for k, v := range count {
		fmt.Printf("%s → %d times\n", k, v)
	}
}
