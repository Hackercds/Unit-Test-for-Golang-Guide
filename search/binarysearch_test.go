package search_test

import (
	"DeepTest/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"empty slice", []int{}, 5, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{3}, 5, -1},
		{"found at start", []int{1, 3, 5, 7, 9}, 1, 0},
		{"found at end", []int{1, 3, 5, 7, 9}, 9, 4},
		{"found in middle", []int{1, 3, 5, 7, 9}, 5, 2},
		{"not found between", []int{1, 3, 5, 7, 9}, 4, -1},
		{"not found less than all", []int{1, 3, 5, 7, 9}, -1, -1},
		{"not found greater than all", []int{1, 3, 5, 7, 9}, 10, -1},
		{"duplicates", []int{1, 3, 3, 3, 7}, 3, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search.BinarySearch(tt.arr, tt.target)
			if got != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.arr, tt.target, got, tt.expected)
			}
		})
	}
}
