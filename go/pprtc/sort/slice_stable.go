package main

import (
	"fmt"
	"sort"
)

type Order struct {
	ID     string
	Price  float64
	Status string
}

func main() {
	orders := []Order{
		{"A100", 300.0, "Pending"},
		{"A101", 120.0, "Shipped"},
		{"A102", 220.0, "Pending"},
		{"A103", 150.0, "Cancelled"},
	}

	sort.SliceStable(orders, func(i, j int) bool {
		return orders[i].Status < orders[j].Status
	})

	fmt.Println("Sorted by Status (stable):")
	for _, o := range orders {
		fmt.Printf("%s | %s | %.2f\n", o.ID, o.Status, o.Price)
	}
}
