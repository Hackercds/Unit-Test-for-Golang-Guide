package search_test

import (
	"DeepTest/search"
	"testing"
)

func BenchmarkBinarySearch(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i * 2
	}
	b.Run("first_element", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search.BinarySearch(arr, 0)
		}
	})
	b.Run("last_element", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search.BinarySearch(arr, 19998)
		}
	})
	b.Run("not_found", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search.BinarySearch(arr, 1)
		}
	})
}

func BenchmarkLinearSearch(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}
	b.Run("first_element", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search.LinearSearch(arr, 0)
		}
	})
	b.Run("last_element", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search.LinearSearch(arr, 9999)
		}
	})
	b.Run("not_found", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			search.LinearSearch(arr, -1)
		}
	})
}
