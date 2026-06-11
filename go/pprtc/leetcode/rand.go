package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"time"
)

func main() {
	// 🔹 Seed for pseudo-random (IMPORTANT)
	// Creates a new independent random generator (local instance)
    // good during concurrency 
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))

	// 1️⃣ Random integer (0 to 99)
	fmt.Println("Random int (0-99):", r.Intn(100))

	// 2️⃣ Random integer in range (10 to 50)
	min, max := 10, 50
	numInRange := r.Intn(max-min+1) + min
	fmt.Println("Random int (10-50):", numInRange)

	// 3️⃣ Random float (0.0 to <1.0)
	fmt.Println("Random float:", r.Float64())

	// 4️⃣ Random boolean
	fmt.Println("Random bool:", r.Intn(2) == 1)

	// 5️⃣ Random element from slice
	arr := []int{10, 20, 30, 40, 50}
	randomElement := arr[r.Intn(len(arr))]
	fmt.Println("Random element from slice:", randomElement)

	// 6️⃣ Shuffle array
	fmt.Println("Before shuffle:", arr)
	r.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println("After shuffle:", arr)

	// 7️⃣ Secure random number (0 to 99)
	secureNum, err := crand.Int(crand.Reader, big.NewInt(100))
	if err != nil {
		fmt.Println("Error generating secure random:", err)
		return
	}
	fmt.Println("Secure random (0-99):", secureNum)

	// 8️⃣ Generate random string (simple example)
	fmt.Println("Random string:", randomString(8, r))
}

// 🔹 Random string generator
func randomString(length int, r *mrand.Rand) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)

	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}

	return string(result)
}

/*


r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
// Creates a new independent random generator (local instance)
// good for concurrency 

rand.Seed(time.Now().UnixNano())
//Modifies the global shared random generator

| Feature        | `rand.Seed()` | `rand.New()`   |
| -------------- | ------------- | -------------- |
| Scope          | Global        | Local instance |
| Concurrency    | ❌ Not safe    | ✅ Safe         |
| Control        | ❌ Shared      | ✅ Independent  |
| Testability    | ❌ Hard        | ✅ Easy         |
| Production use | ❌ Avoid       | ✅ Recommended  |
*/ 