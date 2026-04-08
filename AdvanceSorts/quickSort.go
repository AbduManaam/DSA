package quicksort



func QuickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		QuickSort(arr, low, pivotIndex-1)  // left side
		QuickSort(arr, pivotIndex+1, high) // right side
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // choose last element as pivot
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// place pivot in correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}












