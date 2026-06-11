package make

import "fmt"

func main() {
	A := []int{1, 4, 6, 9, 2}
	sum := 8

	m := make(map[int]int, 0)

	for i := 0; i < len(A); i++ {
		if val, ok := m[sum-A[i]]; ok {
			fmt.Println("Pair Found:", A[i], m[sum-A[i]])
		} else {
			m[A[i]] = i
		}
	}

}
