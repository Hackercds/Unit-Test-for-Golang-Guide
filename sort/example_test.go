package sort_test

import (
	"DeepTest/sort"
	"fmt"
)

func ExampleBubbleSort() {
	arr := []int{3, 1, 2}
	fmt.Println(sort.BubbleSort(arr))
	// Output: [1 2 3]
}

func ExampleQuickSort() {
	arr := []int{5, 2, 8, 1, 9}
	fmt.Println(sort.QuickSort(arr))
	// Output: [1 2 5 8 9]
}

func ExampleMergeSort() {
	arr := []int{4, 2, 7, 1}
	fmt.Println(sort.MergeSort(arr))
	// Output: [1 2 4 7]
}
