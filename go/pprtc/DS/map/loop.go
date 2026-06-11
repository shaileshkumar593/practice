package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initialize map
	freq := map[string]int{
		"apple":  3,
		"banana": 2,
		"cherry": 5,
	}

	fmt.Println("1️⃣ Iterating over key and value (random order):")
	for key, value := range freq {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	fmt.Println("\n2️⃣ Iterating over keys only (random order):")
	for key := range freq {
		fmt.Println("Key only:", key)
	}

	fmt.Println("\n3️⃣ Iterating over values only (random order):")
	for _, value := range freq {
		fmt.Println("Value only:", value)
	}

	// Optional: iterate in sorted key order
	fmt.Println("\n4️⃣ Iterating over map in sorted key order:")
	keys := make([]string, 0, len(freq))
	for key := range freq {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("Key: %s, Value: %d\n", key, freq[key])
	}

	// Optional: iterate by value descending
	fmt.Println("\n5️⃣ Iterating over map sorted by value descending:")
	type kv struct {
		Key   string
		Value int
	}
	var kvs []kv
	for k, v := range freq {
		kvs = append(kvs, kv{k, v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})
	for _, item := range kvs {
		fmt.Printf("Key: %s, Value: %d\n", item.Key, item.Value)
	}
}

/*

	| Operation       | Example                  | Description                               |
| --------------- | ------------------------ | ----------------------------------------- |
| Insert / Update | `m["key"] = value`       | Adds a new key or updates an existing one |
| Delete          | `delete(m, "key")`       | Deletes a key-value pair                  |
| Lookup          | `v := m["key"]`          | Returns value (zero if key missing)       |
| Check existence | `v, ok := m["key"]`      | `ok` is true if key exists                |
| Iterate         | `for k, v := range m {}` | Loop through all key-value pairs          |
| Length          | `len(m)`                 | Returns number of key-value pairs         |
*/
