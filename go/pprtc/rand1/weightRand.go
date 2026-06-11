package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	Name   string
	weight int
}

func weightedRandom(items []Item) string {
	total := 0

	for _, item := range items {
		total += item.weight
	}

	r := rand.Intn(total)
	current := 0

	for _, item := range items {
		current += item.weight

		if r < current {
			return item.Name
		}
	}

	return ""
}

func main() {
	rand.Seed(time.Now().UnixNano())

	items := []Item{
		{"India", 50},
		{"USA", 30},
		{"UK", 20},
	}

	fmt.Println(weightedRandom(items))
}
