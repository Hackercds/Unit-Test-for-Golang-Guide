package mathutil_test

import (
	"DeepTest/mathutil"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"both zero", 0, 0, 0},
		{"one zero", 12, 0, 12},
		{"one zero reversed", 0, 8, 8},
		{"normal", 12, 8, 4},
		{"coprime", 17, 5, 1},
		{"same numbers", 7, 7, 7},
		{"negative a", -12, 8, 4},
		{"negative b", 12, -8, 4},
		{"both negative", -12, -8, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mathutil.GCD(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
