package mathutil_test

import (
	"DeepTest/mathutil"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"negative", -1, -1},
		{"0!", 0, 1},
		{"1!", 1, 1},
		{"5!", 5, 120},
		{"10!", 10, 3628800},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mathutil.Factorial(tt.input)
			if got != tt.expected {
				t.Errorf("Factorial(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
