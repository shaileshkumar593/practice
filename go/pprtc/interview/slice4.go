package main

import (
	"fmt"
	"sort"
)

// 🔹 71. Transpose of a matrix
func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}

// 🔹 72. Matrix addition
func matrixAdd(a, b [][]int) [][]int {
	rows, cols := len(a), len(a[0])
	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

// 🔹 73. Matrix multiplication
func matrixMul(a, b [][]int) [][]int {
	rows, cols, n := len(a), len(b[0]), len(b)
	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			for k := 0; k < n; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

// 🔹 74. Identity matrix
func identity(n int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		mat[i][i] = 1
	}
	return mat
}

// 🔹 75. Flatten + unique
func flattenUnique(matrix [][]int) []int {
	seen := map[int]bool{}
	var res []int
	for _, row := range matrix {
		for _, v := range row {
			if !seen[v] {
				res = append(res, v)
				seen[v] = true
			}
		}
	}
	return res
}

// 🔹 76. Sliding window max
func slidingMax(nums []int, k int) []int {
	var res []int
	for i := 0; i <= len(nums)-k; i++ {
		max := nums[i]
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res
}

// 🔹 77. Prefix min
func prefixMin(nums []int) []int {
	mins := make([]int, len(nums))
	min := nums[0]
	for i, v := range nums {
		if v < min {
			min = v
		}
		mins[i] = min
