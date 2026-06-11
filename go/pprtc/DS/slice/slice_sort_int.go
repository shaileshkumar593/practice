package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// -------------------------------
	// 1. sort.Ints
	// -------------------------------
	nums := []int{5, 2, 9, 1, 5, 6}
	sort.Ints(nums)
	fmt.Println("sort.Ints:", nums) // Ascending integers

	// -------------------------------
	// 2. sort.Strings
	// -------------------------------
	words := []string{"banana", "apple", "cherry", "apple"}
	sort.Strings(words)
	fmt.Println("sort.Strings:", words) // Ascending strings

	// -------------------------------
	// 3. sort.Float64s
	// -------------------------------
	floats := []float64{3.14, 2.71, 1.41, 2.71}
	sort.Float64s(floats)
	fmt.Println("sort.Float64s:", floats)

	// -------------------------------
	// 4. sort.Sort (custom)
	// -------------------------------
	persons := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
	}
	// Sort by Age ascending
	sort.Sort(ByAge(persons))
	fmt.Println("sort.Sort (by age):", persons)

	// -------------------------------
	// 5. sort.Slice
	// -------------------------------
	meetings := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
	}
	// Sort by Age descending
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].Age > meetings[j].Age
	})
	fmt.Println("sort.Slice (Age desc):", meetings)

	// -------------------------------
	// 6. sort.SliceStable
	// -------------------------------
	stableMeetings := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
	}
	// Sort by Age ascending but stable
	sort.SliceStable(stableMeetings, func(i, j int) bool {
		return stableMeetings[i].Age < stableMeetings[j].Age
	})
	fmt.Println("sort.SliceStable (Age asc, stable):", stableMeetings)

	// -------------------------------
	// 7. sort.Stable
	// -------------------------------
	morePersons := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
	}
	sort.Stable(ByAge(morePersons)) // Stable sort
	fmt.Println("sort.Stable (by age, stable):", morePersons)
}

// Custom type for sort.Sort
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

/*

	✅ Explanation:
Function				Example in code								Notes
sort.Ints				sort.Ints(nums)								Simple integer ascending sort
sort.Strings			sort.Strings(words)							Simple string ascending sort
sort.Float64s			sort.Float64s(floats)						Simple float ascending sort
sort.Sort				sort.Sort(ByAge(persons))					Custom type, implements sort.Interface
sort.Slice				sort.Slice(meetings, func...)				Sort any slice using a lambda
sort.SliceStable		sort.SliceStable(stableMeetings, func...)	Stable sort preserves relative order for equal elements
sort.Stable				sort.Stable(morePersons)					Sort any slice implementing sort.Interface stably


Key Difference:

Introsort (sort.Ints, sort.Slice, etc.): Fast, not stable. Equal elements may reorder.

MergeSort-like (sort.SliceStable, sort.Stable): Stable, preserves order of equal elements.

sort.SliceStable sorts a slice while preserving the relative order of equal elements.

It is stable, unlike sort.Slice which can reorder elements with equal keys.

Useful when you have multiple keys or want to keep previous ordering intact for duplicates.


Function			Algorithm		Stable			Complexity
sort.Ints			Introsort		No				Avg O(n log n)
sort.Strings		Introsort		No				Avg O(n log n)
sort.Float64s		Introsort		No				Avg O(n log n)
sort.Sort			Introsort		No				Avg O(n log n)
sort.Slice			Introsort		No				Avg O(n log n)
sort.SliceStable	MergeSort-like	Yes				O(n log n)
sort.Stable			MergeSort-like	Yes				O(n log n)


| Operation             | Go Syntax / Function                        |
| --------------------- | ------------------------------------------- |
| Create                | `var s []int`, `s := []int{1,2}`            |
| Access                | `s[0]`                                      |
| Length & Capacity     | `len(s)`, `cap(s)`                          |
| Slice Subset          | `s[1:4]`, `s[:3]`, `s[2:]`                  |
| Append                | `append(s, 3)`                              |
| Delete                | `s = append(s[:i], s[i+1:]...)`             |
| Copy                  | `copy(dst, src)`                            |
| Iterate               | `for i, v := range s {}`                    |
| Sort                  | `sort.Ints(s)`, `sort.Slice(s, lessFunc)`   |
| Reverse               | `sort.Sort(sort.Reverse(sort.IntSlice(s)))` |
| Search (sorted slice) | `sort.SearchInts(s, value)`                 |

*/
