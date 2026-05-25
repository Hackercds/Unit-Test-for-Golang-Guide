package sort_test

import (
	"DeepTest/sort"
	"math/rand"
	"testing"
)

func generateRandomSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(10000)
	}
	return s
}

func BenchmarkBubbleSort(b *testing.B) {
	sizes := []int{10, 100, 500}
	for _, size := range sizes {
		original := generateRandomSlice(size)
		b.Run(sizeName(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := make([]int, size)
				copy(arr, original)
				b.StartTimer()
				sort.BubbleSort(arr)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	sizes := []int{10, 100, 1000}
	for _, size := range sizes {
		original := generateRandomSlice(size)
		b.Run(sizeName(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := make([]int, size)
				copy(arr, original)
				b.StartTimer()
				sort.QuickSort(arr)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	sizes := []int{10, 100, 1000}
	for _, size := range sizes {
		original := generateRandomSlice(size)
		b.Run(sizeName(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.MergeSort(original)
			}
		})
	}
}

func sizeName(n int) string {
	switch n {
	case 10:
		return "size=10"
	case 100:
		return "size=100"
	case 500:
		return "size=500"
	case 1000:
		return "size=1000"
	default:
		return "unknown"
	}
}
