package mathutil_test

import (
	"DeepTest/mathutil"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"negative", -1, -1},
		{"F(0)", 0, 0},
		{"F(1)", 1, 1},
		{"F(2)", 2, 1},
		{"F(10)", 10, 55},
		{"F(20)", 20, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mathutil.Fibonacci(tt.input)
			if got != tt.expected {
				t.Errorf("Fibonacci(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
