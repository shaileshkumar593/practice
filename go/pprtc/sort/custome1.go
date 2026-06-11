package main

import (
	"fmt"
	"sort"
)

// Person represents a person with a name and age.
type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for sorting a slice of Person by Age.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age } // Ascending order by Age

// ByName implements sort.Interface for sorting a slice of Person by Name.
type ByName []Person

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name } // Ascending order by Name

func main() {
	// Create a slice of Person structs.
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 25}, // Another person with the same age
	}

	fmt.Println("Original people:", people)

	// Sort by Age (ascending)
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by Age:", people)

	// Sort by Name (ascending)
	sort.Sort(ByName(people))
	fmt.Println("Sorted by Name:", people)

	// Custom sorting with a less function (e.g., descending age)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age // Descending order by Age
	})
	fmt.Println("Sorted by Age (descending) using sort.Slice:", people)

	// Custom sorting with multiple criteria (e.g., primary by age ascending, secondary by name ascending)
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age != people[j].Age {
			return people[i].Age < people[j].Age // Primary sort by Age ascending
		}
		return people[i].Name < people[j].Name // Secondary sort by Name ascending
	})
	fmt.Println("Sorted by Age (asc) then Name (asc) using sort.Slice:", people)
}
