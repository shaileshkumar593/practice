package main

import "fmt"

/*
	we need to declare a variable or parameter but don't actually need to use it.
	In such cases, we can use an empty struct as a placeholder to avoid wasting memory.
	For example, we can use map[string]struct{} to represent a collection of key-value pairs,
	where an empty struct can be used to represent a case where we only need the key and not the value.

	we use an empty struct as the value of map[string]struct{} to indicate that we only need the key and not the value.
	This can avoid wasting memory and improve the readability of the code.

*/

func main() {
	m := map[string]struct{}{
		"apple":    {},
		"banana":   {},
		"cherry":   {},
		"durian":   {},
		"elder":    {},
		"fig":      {},
		"grape":    {},
		"honeydew": {},
	}

	for k, _ := range m {
		fmt.Println(k)
	}
}
