package main

import "fmt"

// CASE 1: Named return + defer modifies it
func case1() (i int) {
	i = 10

	defer func() {
		i = i + 5
		// WHY: named return variable is modified AFTER return is prepared
	}()

	return
	// FINAL OUTPUT: 15
}

// CASE 2: Named return + explicit return value
func case2() (i int) {
	defer func() {
		i = 100
		// WHY: defer runs AFTER return value (50) is assigned to i
	}()

	return 50
	// FINAL OUTPUT: 100
}

// CASE 3: Anonymous return (defer does NOT affect return)
func case3() int {
	i := 10

	defer func() {
		i = 100
		// WHY: modifies local variable, NOT return value (copy already made)
	}()

	return i
	// FINAL OUTPUT: 10
}

// CASE 4: return expression + defer increment
func case4() int {
	i := 10

	defer func() {
		i++
		// WHY: return already copied value before defer runs
	}()

	return i
	// FINAL OUTPUT: 10
}

// CASE 5: Named return + return expression + defer
func case5() (i int) {
	i = 10

	defer func() {
		i++
		// WHY: return value is stored in i first, then defer modifies it
	}()

	return 20
	// FINAL OUTPUT: 21
}

// CASE 6: Multiple defer (stack order)
func case6() (i int) {
	i = 1

	defer func() {
		i++
		// runs LAST (LIFO stack)
	}()

	defer func() {
		i = i * 10
		// runs FIRST
	}()

	return 5

	/*
		STEP-BY-STEP:
		i = 5
		defer2 → i = 50
		defer1 → i = 51

		FINAL OUTPUT: 51
	*/
}

func main() {
	fmt.Println("Case 1:", case1()) // 15
	fmt.Println("Case 2:", case2()) // 100
	fmt.Println("Case 3:", case3()) // 10
	fmt.Println("Case 4:", case4()) // 10
	fmt.Println("Case 5:", case5()) // 21
	fmt.Println("Case 6:", case6()) // 51
}
