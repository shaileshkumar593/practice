package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scanf("%d", &n)
	res := n
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if i <= res {
				res = i
			}
		}
	}
	fmt.Printf("The smallest divisor of the number is: %d", res)
}
