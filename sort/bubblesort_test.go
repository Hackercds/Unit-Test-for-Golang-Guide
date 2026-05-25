package sort_test

import (
	"DeepTest/sort"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"with duplicates", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},
		{"negative numbers", []int{-3, 0, 5, -1, 2}, []int{-3, -1, 0, 2, 5}},
		{"all same value", []int{7, 7, 7, 7}, []int{7, 7, 7, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := copySlice(tt.input)
			got := sort.BubbleSort(input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func copySlice(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	return c
}
