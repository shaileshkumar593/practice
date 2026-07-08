package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Expression struct {
	coefficients map[rune]int
	constant     int
}

// Parse expression
func parseExpression(expr string) Expression {

	expr = strings.ReplaceAll(expr, " ", "")

	result := Expression{
		coefficients: make(map[rune]int),
	}

	sign := 1
	number := ""

	for i := 0; i < len(expr); i++ {

		ch := rune(expr[i])

		// Digit
		if unicode.IsDigit(ch) {
			number += string(ch)
			continue
		}

		// + or -
		if ch == '+' || ch == '-' {

			// Previous number is constant
			if number != "" {
				value, _ := strconv.Atoi(number)
				result.constant += sign * value
				number = ""
			}

			if ch == '+' {
				sign = 1
			} else {
				sign = -1
			}

			continue
		}

		// Variable
		if unicode.IsLetter(ch) {

			coef := 1

			if number != "" {
				coef, _ = strconv.Atoi(number)
			}

			result.coefficients[ch] += sign * coef

			number = ""
		}
	}

	// Last constant
	if number != "" {
		value, _ := strconv.Atoi(number)
		result.constant += sign * value
	}

	return result
}

// Add two expressions
func addExpression(e1, e2 Expression) Expression {

	result := Expression{
		coefficients: make(map[rune]int),
	}

	for k, v := range e1.coefficients {
		result.coefficients[k] = v
	}

	for k, v := range e2.coefficients {
		result.coefficients[k] += v
	}

	result.constant = e1.constant + e2.constant

	return result
}

// Print expression
func printExpression(e Expression) {

	first := true

	for variable, coefficient := range e.coefficients {

		if coefficient == 0 {
			continue
		}

		if !first && coefficient > 0 {
			fmt.Print("+")
		}

		if coefficient == 1 {
			fmt.Printf("%c", variable)
		} else if coefficient == -1 {
			fmt.Printf("-%c", variable)
		} else {
			fmt.Printf("%d%c", coefficient, variable)
		}

		first = false
	}

	if e.constant > 0 {
		fmt.Printf("+%d", e.constant)
	} else if e.constant < 0 {
		fmt.Printf("%d", e.constant)
	}

	fmt.Println()
}

func main() {

	exp1 := "5x+27y+10"
	exp2 := "2x-y+2"

	e1 := parseExpression(exp1)
	e2 := parseExpression(exp2)

	fmt.Println("Expression 1 Map:")
	fmt.Println(e1.coefficients)
	fmt.Println("Constant:", e1.constant)

	fmt.Println()

	fmt.Println("Expression 2 Map:")
	fmt.Println(e2.coefficients)
	fmt.Println("Constant:", e2.constant)

	fmt.Println()

	result := addExpression(e1, e2)

	fmt.Println("Result Map:")
	fmt.Println(result.coefficients)
	fmt.Println("Constant:", result.constant)

	fmt.Println()

	fmt.Print("Final Expression: ")
	printExpression(result)
}