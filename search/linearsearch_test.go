package search_test

import (
	"DeepTest/search"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"empty slice", []int{}, 5, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{3}, 5, -1},
		{"first occurrence", []int{3, 1, 3, 5}, 3, 0},
		{"not found", []int{1, 2, 3}, 4, -1},
		{"negative numbers", []int{-5, 0, 3}, -5, 0},
		{"unsorted works", []int{9, 2, 7, 1}, 7, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search.LinearSearch(tt.arr, tt.target)
			if got != tt.expected {
				t.Errorf("LinearSearch(%v, %d) = %d, want %d", tt.arr, tt.target, got, tt.expected)
			}
		})
	}
}
