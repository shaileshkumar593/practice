package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{-1, -1},
	}

	for _, tt := range tests {
		if res := factorial(tt.input); res != tt.expected {
			t.Errorf("got %d, want %d", res, tt.expected)
		}
	}
}
