package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Country struct {
	Name       string
	Percentage int
}

func pickCountry(countries []Country) string {
	total := 0
	for _, c := range countries {
		total += c.Percentage
	}
	fmt.Println("total val ", total)
	r := rand.Intn(total) // random number in [0, total)

	current := 0
	for _, c := range countries {
		current += c.Percentage
		fmt.Println("-------", current)
		if r < current {
			return c.Name
		}
	}
	return ""
}

func main() {
	rand.Seed(time.Now().UnixNano())

	countries := []Country{
		{"India", 50},
		{"USA", 30},
		{"UK", 20},
	}

	count := map[string]int{}

	// Run multiple times to see distribution
	for i := 0; i < 100000; i++ {
		c := pickCountry(countries)
		count[c]++
	}

	for k, v := range count {
		fmt.Printf("%s → %d times\n", k, v)
	}
}
