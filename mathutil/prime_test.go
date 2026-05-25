package mathutil_test

import (
	"DeepTest/mathutil"
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"negative", -5, false},
		{"zero", 0, false},
		{"one", 1, false},
		{"two", 2, true},
		{"three", 3, true},
		{"four", 4, false},
		{"even >2", 100, false},
		{"prime 17", 17, true},
		{"large prime", 9973, true},
		{"large non-prime", 9999, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mathutil.IsPrime(tt.input)
			if got != tt.expected {
				t.Errorf("IsPrime(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
