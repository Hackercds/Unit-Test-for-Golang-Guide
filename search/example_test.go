package search_test

import (
	"DeepTest/search"
	"fmt"
)

func ExampleBinarySearch() {
	arr := []int{1, 3, 5, 7, 9}
	fmt.Println(search.BinarySearch(arr, 7))
	fmt.Println(search.BinarySearch(arr, 4))
	// Output:
	// 3
	// -1
}

func ExampleLinearSearch() {
	arr := []int{4, 2, 8, 1}
	fmt.Println(search.LinearSearch(arr, 8))
	// Output: 2
}
