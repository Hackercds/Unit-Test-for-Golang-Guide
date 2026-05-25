package search

// LinearSearch performs linear search on a slice.
// Returns the index of the first occurrence of target, or -1 if not found.
func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
