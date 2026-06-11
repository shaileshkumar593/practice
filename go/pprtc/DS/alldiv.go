package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter an integer:")
	fmt.Scanf("%d", &n)
	fmt.Println("The divisors of the number are:")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
