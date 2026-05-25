package main

import (
	"fmt"

	"DeepTest/mathutil"
	"DeepTest/search"
	"DeepTest/sort"
	"DeepTest/tttys"
	"DeepTest/utils"
)

func main() {
	fmt.Println("=== DeepTest Algorithm Demo ===")

	// Existing utilities
	fmt.Println(tttys.Greeting())
	fmt.Println(tttys.Initialize())
	fmt.Println(utils.ConfigMessage())

	// Sorting demo
	fmt.Println("\n--- Sorting ---")
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original: %v\n", arr)
	fmt.Printf("BubbleSort: %v\n", sort.BubbleSort(copySlice(arr)))
	fmt.Printf("QuickSort: %v\n", sort.QuickSort(copySlice(arr)))
	fmt.Printf("MergeSort: %v\n", sort.MergeSort(arr))
	fmt.Printf("Original (unchanged by MergeSort): %v\n", arr)

	// Searching demo
	fmt.Println("\n--- Searching ---")
	sorted := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Printf("Sorted array: %v\n", sorted)
	fmt.Printf("BinarySearch 7: %d\n", search.BinarySearch(sorted, 7))
	fmt.Printf("BinarySearch 4: %d\n", search.BinarySearch(sorted, 4))
	fmt.Printf("LinearSearch 7: %d\n", search.LinearSearch(sorted, 7))

	// Math demo
	fmt.Println("\n--- Math ---")
	fmt.Printf("Fibonacci(10): %d\n", mathutil.Fibonacci(10))
	fmt.Printf("GCD(48, 18): %d\n", mathutil.GCD(48, 18))
	fmt.Printf("IsPrime(17): %v\n", mathutil.IsPrime(17))
	fmt.Printf("IsPrime(100): %v\n", mathutil.IsPrime(100))
	fmt.Printf("Factorial(5): %d\n", mathutil.Factorial(5))
}

func copySlice(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	return c
}
