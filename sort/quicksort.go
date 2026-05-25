package sort

// QuickSort sorts arr in ascending order in-place using quick sort.
// Returns the same slice for convenience.
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	quickSortRange(arr, 0, len(arr)-1)
	return arr
}

func quickSortRange(arr []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(arr, low, high)
	quickSortRange(arr, low, p-1)
	quickSortRange(arr, p+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
