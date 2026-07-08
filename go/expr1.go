package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func parse(expr string, mp map[string][]int, constant *[]int) {

	expr = strings.ReplaceAll(expr, " ", "")

	sign := 1
	number := ""

	for i := 0; i < len(expr); i++ {

		ch := rune(expr[i])

		// Number
		if unicode.IsDigit(ch) {
			number += string(ch)
			continue
		}

		// Variable
		if unicode.IsLetter(ch) {

			coef := 1

			if number != "" {
				coef, _ = strconv.Atoi(number)
			}

			coef *= sign

			mp[string(ch)] = append(mp[string(ch)], coef)

			number = ""

			continue
		}

		// + or -
		if ch == '+' || ch == '-' {

			// Remaining number is constant
			if number != "" {
				value, _ := strconv.Atoi(number)
				*constant = append(*constant, sign*value)
				number = ""
			}

			if ch == '+' {
				sign = 1
			} else {
				sign = -1
			}
		}
	}

	// Last constant
	if number != "" {
		value, _ := strconv.Atoi(number)
		*constant = append(*constant, sign*value)
	}
}

func evaluate(mp map[string][]int) map[string]int {

	result := make(map[string]int)

	for variable, coeffs := range mp {

		sum := 0

		for _, coefficient := range coeffs {
			sum += coefficient
		}

		result[variable] = sum
	}

	return result
}

func sumConstant(arr []int) int {

	sum := 0

	for _, value := range arr {
		sum += value
	}

	return sum
}

func main() {

	exp1 := "5x+27y+10"
	exp2 := "2x-y+2"

	mp := make(map[string][]int)

	constants := []int{}

	parse(exp1, mp, &constants)
	parse(exp2, mp, &constants)

	fmt.Println("Map After Parsing")
	fmt.Println("-----------------")

	for key, value := range mp {
		fmt.Println(key, "=>", value)
	}

	fmt.Println()

	fmt.Println("Constants =>", constants)

	fmt.Println()

	result := evaluate(mp)

	fmt.Println("Final Coefficients")
	fmt.Println("------------------")

	for key, value := range result {
		fmt.Println(key, "=", value)
	}

	fmt.Println()

	fmt.Println("Constant =", sumConstant(constants))

	fmt.Println()

	fmt.Printf("Final Expression = %dx + %dy + %d\n",
		result["x"],
		result["y"],
		sumConstant(constants))
}